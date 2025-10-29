package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Nova diretiva 'embed' para incluir todos os arquivos da pasta 'public' no binário.
// A variável 'arquivosPublicos' representará esse conjunto de arquivos.

/*
go:embed -> Atenção! A próxima declaração de variável
terá seu conteúdo preenchido com arquivos do disco.
all:public -> É o caminho para o arquivo ou pasta que você quer embutir (pasta public).
O caminho é relativo ao arquivo .go onde a diretiva está. O all indica que sao todos os arquivos.
Esta diretiva diz respeito ao arquivosPublicos  logo abaixo
*/

//go:embed all:public
var arquivosPublicos embed.FS

// Nova constante para o nome do arquivo de configuração, facilitando a manutenção.
const nomeArquivoConfig = "config.yaml"

// Structs que vao "entender o formato das respostas JSON"

// Estrutura para ler e escrever a configuração do servidor.
// As 'tags' `mapstructure` e `yaml` ajudam as bibliotecas a ler/escrever os dados corretamente.
type Configuracao struct {
	Servidor struct {
		Host  string `mapstructure:"host" yaml:"host"`
		Porta string `mapstructure:"porta" yaml:"porta"`
	} `mapstructure:"servidor" yaml:"servidor"`
}

// Estrutura para a nossa resposta da API
type MensagemAPI struct {
	Texto     string `json:"texto"`
	Timestamp string `json:"timestamp"`
	Frase     string `json:"frase"`
}

// Estrutura para decodificar a resposta da API CoinGecko
type BitCoinDados struct {
	Bitcoin struct {
		BRL float64 `json:"brl"`
	} `json:"bitcoin"`
}

// Estrutura para decodificar a resposta da API wttr.in
type ClimaDados struct {
	CondicoesTempo []struct {
		TempC    string `json:"temp_C"`
		DescTemp []struct {
			Valor string `json:"value"`
		} `json:"weatherDesc"`
	} `json:"current_condition"`
}

// Lista de frases motivacionais que nosso backend pode escolher.
var frasesMotivacionais = []string{
	"O sucesso é a soma de pequenos esforços repetidos dia após dia.",
	"Acredite em você mesmo e tudo será possível.",
	"O único lugar onde o sucesso vem antes do trabalho é no dicionário.",
	"Comece onde você está. Use o que você tem. Faça o que você pode.",
}

// Teremos apenas um handler, que faz tudo.
func mensagemHandler(w http.ResponseWriter, r *http.Request) {
	// Configura os cabeçalhos essenciais (CORS e tipo de conteúdo)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// --- Lógica para escolher a frase aleatória (movida para cá) ---
	indiceAleatorio := rand.Intn(len(frasesMotivacionais))
	fraseEscolhida := frasesMotivacionais[indiceAleatorio]

	// --- Monta a resposta completa com todos os dados ---
	respostaCompleta := MensagemAPI{
		Texto:     "Dados recebidos do Backend Go! 🚀",
		Timestamp: time.Now().Format(time.RFC3339),
		Frase:     fraseEscolhida, // Adiciona a frase escolhida à resposta
	}

	// Codifica e envia a resposta completa
	err := json.NewEncoder(w).Encode(respostaCompleta)
	if err != nil {
		log.Printf("Erro ao codificar JSON: %v", err)
	}
}

// Handler para cotação do Bitcoin
func bitcoinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// 1. Chama a API externa da CoinGecko
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=brl")
	if err != nil {
		http.Error(w, "Falha ao buscar dados da CoinGecko", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 2. Decodifica a resposta JSON na nossa struct
	var bitcoinDado BitCoinDados
	if err := json.NewDecoder(resp.Body).Decode(&bitcoinDado); err != nil {
		http.Error(w, "Falha ao decodificar resposta da CoinGecko", http.StatusInternalServerError)
		return
	}

	// 3. Envia os dados para o frontend
	json.NewEncoder(w).Encode(bitcoinDado)
}

// Handler para a temperatura
func climaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// IMPORTANTE: Troque "Joao+Pessoa" pela sua cidade!
	// O formato é Nome+Da+Cidade, sem acentos.
	cidade := "Joao+Pessoa"
	url := fmt.Sprintf("http://wttr.in/%s?format=j1", cidade)

	// 1. Chama a API externa do wttr.in
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERRO: Falha ao fazer a requisição para wttr.in: %v", err)
		http.Error(w, "Falha ao buscar dados de temperatura", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERRO: Falha ao ler o corpo da resposta de wttr.in: %v", err)
		http.Error(w, "Falha ao ler a resposta da API de clima", http.StatusInternalServerError)
		return
	}
	//log.Printf("Resposta recebida de wttr.in: %s", string(bodyBytes))

	// 2. Decodifica a resposta JSON
	var climadado ClimaDados
	if err := json.Unmarshal(bodyBytes, &climadado); err != nil {
		log.Printf("ERRO: Falha ao decodificar o JSON de wttr.in: %v", err)
		http.Error(w, "Falha ao decodificar resposta de temperatura", http.StatusInternalServerError)
		return
	}

	if len(climadado.CondicoesTempo) == 0 {
		log.Println("AVISO: O JSON foi decodificado, mas o array 'current_condition' está vazio.")
		http.Error(w, "Dados de condição atual não encontrados na resposta da API", http.StatusInternalServerError)
		return
	}

	// 3. Envia os dados para o frontend
	json.NewEncoder(w).Encode(climadado)
}

// Nova função que usa o Viper para carregar o arquivo de configuração que já existe.
func carregarConfiguracao() (config Configuracao, err error) {
	// Define o nome do arquivo que o Viper deve procurar.
	viper.SetConfigFile(nomeArquivoConfig)
	// Permite que variáveis de ambiente sobreescrevam as do arquivo.
	viper.AutomaticEnv()

	// Tenta ler o arquivo de configuração.
	if err = viper.ReadInConfig(); err != nil {
		// Retorna um erro se não conseguir ler o arquivo.
		return
	}

	// Transfere os dados lidos do arquivo para a nossa struct 'Configuracao'.
	err = viper.Unmarshal(&config)
	return
}

// Função para o prompt interativo que cria o arquivo de configuração na primeira execução.
func perguntarECriarConfiguracao() (Configuracao, error) {
	// Declara uma struct anônima para receber as respostas das perguntas.
	// As tags `survey:"porta"` correspondem ao nome que daremos a cada pergunta.
	respostas := struct {
		Porta string `survey:"porta"`
		Host  string `survey:"host"`
	}{}

	// Cria o conjunto de perguntas que serão feitas ao usuário.
	var perguntas = []*survey.Question{
		{
			Name: "porta",
			Prompt: &survey.Input{
				Message: "Qual porta o servidor deve usar?",
				Default: "8080", // Podemos sugerir um valor padrão.
			},
			Validate: survey.Required, // Torna a resposta obrigatória.
		},
		{
			Name: "host",
			Prompt: &survey.Select{ // Usamos um seletor para guiar o usuário!
				Message: "Qual endereço o servidor deve ouvir?",
				Options: []string{"0.0.0.0 (recomendado, acesso local e de rede)", "localhost (apenas acesso local)"},
				Default: "0.0.0.0 (recomendado, acesso local e de rede)",
			},
		},
	}

	// Imprime o cabeçalho da configuração.
	fmt.Println("--- Configuração Inicial ---")
	fmt.Printf("Arquivo '%s' não encontrado. Vamos criá-lo.\n", nomeArquivoConfig)

	// Faz a "mágica": exibe as perguntas e preenche a struct 'respostas' com o que o usuário digitar.
	err := survey.Ask(perguntas, &respostas)
	if err != nil {
		// Se o usuário cancelar (Ctrl+C), o erro é tratado aqui.
		return Configuracao{}, err
	}

	// Cria a nossa struct de configuração final a partir das respostas.
	var configPreenchida Configuracao
	configPreenchida.Servidor.Porta = respostas.Porta
	// O 'survey.Select' retorna o texto completo, então pegamos só o que importa (antes do espaço).
	fmt.Sscanf(respostas.Host, "%s", &configPreenchida.Servidor.Host)

	// Converte a struct para o formato de texto YAML.
	dadosYaml, err := yaml.Marshal(&configPreenchida)
	if err != nil {
		return Configuracao{}, fmt.Errorf("erro ao gerar YAML: %w", err)
	}

	// Escreve os dados YAML no arquivo 'config.yaml'.
	err = os.WriteFile(nomeArquivoConfig, dadosYaml, 0644)
	if err != nil {
		return Configuracao{}, fmt.Errorf("erro ao escrever arquivo de configuração: %w", err)
	}

	fmt.Printf("\n✅ Arquivo '%s' criado com sucesso!\n", nomeArquivoConfig)
	return configPreenchida, nil
}

// Nova função auxiliar que verifica se um arquivo existe no disco.
func arquivoExiste(nomeArquivo string) bool {
	// Tenta obter informações sobre o arquivo.
	info, err := os.Stat(nomeArquivo)
	// Se o erro for do tipo "Não Existe", retorna falso.
	if os.IsNotExist(err) {
		return false
	}
	// Garante que o caminho não é um diretório, mas sim um arquivo.
	return !info.IsDir()
}

func main() {

	// Declara as variáveis que usaremos para a configuração e para erros.
	var config Configuracao
	var err error

	// LÓGICA PRINCIPAL DA INICIALIZAÇÃO:
	// Verifica se o arquivo de configuração já existe.
	if arquivoExiste(nomeArquivoConfig) {
		// Se existir, carrega as configurações dele.
		log.Printf("Carregando configuração de '%s'...", nomeArquivoConfig)
		config, err = carregarConfiguracao()
	} else {
		// Se não existir, inicia o processo interativo para criá-lo.
		config, err = perguntarECriarConfiguracao()
	}

	// Se houve qualquer erro no processo de configuração, encerra o programa.
	if err != nil {
		log.Fatalf("Processo de configuração cancelado ou falhou: %v", err)
	}

	// Prepara o sistema de arquivos embutido para ser servido via HTTP.
	// O 'fs.Sub' cria uma "visão" da pasta 'public' a partir da raiz dos arquivos embutidos.
	arquivosEstaticosFS, err := fs.Sub(arquivosPublicos, "public")
	if err != nil {
		log.Fatal(err)
	}

	//log.Println("--- Verificando arquivos embutidos (embed)... ---")
	//err = fs.WalkDir(arquivosPublicos, ".", func(path string, d fs.DirEntry, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//	log.Printf("Arquivo encontrado: %s\n", path)
	//	return nil
	//})
	//if err != nil {
	//	log.Fatalf("Erro ao listar arquivos embutidos: %v", err)
	//}
	//log.Println("--- Fim da verificação ---")

	// Cria um novo roteador (multiplexer) para registrar nossas rotas.
	mux := http.NewServeMux()

	// Registra nossas rotas de API no roteador.
	mux.HandleFunc("/api/mensagem", mensagemHandler)
	mux.HandleFunc("/api/bitcoin", bitcoinHandler)
	mux.HandleFunc("/api/clima", climaHandler)

	// Registra o manipulador de arquivos estáticos para a rota raiz "/".
	// Qualquer rota que não for de API será tratada por ele, servindo o frontend Svelte.
	mux.Handle("/", http.FileServer(http.FS(arquivosEstaticosFS)))

	// Monta o endereço completo do servidor (host:porta) a partir da configuração carregada.
	enderecoServidor := fmt.Sprintf("%s:%s", config.Servidor.Host, config.Servidor.Porta)

	// Para criar links clicáveis, se o host for 0.0.0.0, usamos 'localhost'.
	hostParaLink := config.Servidor.Host
	if hostParaLink == "0.0.0.0" {
		hostParaLink = "localhost"
	}
	baseUrl := fmt.Sprintf("http://%s:%s", hostParaLink, config.Servidor.Porta)

	// Exibe mensagens informativas no console sobre o estado do servidor.
	fmt.Printf("\n🚀 Servidor Go rodando em %s\n", baseUrl)
	fmt.Println("   - Frontend Svelte servido na rota '/'")
	fmt.Println("   - API disponível nas rotas '/api/*'")

	fmt.Println("Endpoints de API disponíveis:")
	fmt.Printf("  - %s/api/mensagem\n", baseUrl)
	fmt.Printf("  - %s/api/bitcoin\n", baseUrl)
	fmt.Printf("  - %s/api/clima\n", baseUrl)

	// Inicia o servidor HTTP e o faz "ouvir" no endereço configurado.
	// O 'log.Fatal' fará com que o programa encerre se houver um erro ao iniciar o servidor.
	log.Fatal(http.ListenAndServe(enderecoServidor, mux))
}
