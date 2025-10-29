package main

import (
	"bufio"          // Para ler a entrada do usuário de forma eficiente.
	"database/sql"   // Fornece a interface padrão para bancos de dados SQL.
	"fmt"            // Pacote para formatação de entrada e saída, como imprimir no console.
	"log"            // Usado para registrar mensagens de erro fatais.
	"os"             // Fornece funções para interagir com o sistema operacional (arquivos, argumentos, etc.).
	"strconv"        // Para conversão entre strings e outros tipos (ex: string para inteiro).
	"strings"        // Funções para manipulação de strings.
	"text/tabwriter" // Pacote para criar tabelas bem alinhadas no console.

	// O '_' significa que estamos importando o pacote por seus efeitos colaterais,
	// que neste caso é registrar o driver do SQLite. Não usamos o pacote diretamente.
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra" // Biblioteca para criar aplicações de linha de comando (CLI).
	"gopkg.in/yaml.v3"       // Pacote para trabalhar com arquivos de configuração no formato YAML.
)

// ============== ESTRUTURAS DE DADOS ==============

// Configuracao define a estrutura do nosso arquivo config.yml.
// A tag `yaml:"..."` mapeia o campo da struct para a chave no arquivo YAML.
type Configuracao struct {
	CaminhoBancoDados string `yaml:"database_path"`
}

// ============== VARIÁVEIS GLOBAIS ==============

var (
	// arquivoConfig define o nome padrão do nosso arquivo de configuração.
	arquivoConfig = "config.yml"
	// bancoDeDados é a variável global que manterá a conexão com o banco de dados ativa.
	bancoDeDados *sql.DB
)

// ============== LÓGICA DE CONFIGURAÇÃO (config.yml) ==============

// salvarConfiguracao serializa (converte) a struct Configuracao para o formato YAML
// e a salva no arquivo definido pela variável 'arquivoConfig'.
func salvarConfiguracao(config Configuracao) error {
	// Converte a struct para bytes no formato YAML.
	dados, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	// Escreve os bytes no arquivo com permissões de leitura/escrita para o dono.
	return os.WriteFile(arquivoConfig, dados, 0644)
}

// carregarConfiguracao lê e interpreta o arquivo config.yml.
// Se o arquivo não existir, inicia um setup interativo para criá-lo.
func carregarConfiguracao() (Configuracao, error) {
	var config Configuracao

	// Verifica se o arquivo de configuração existe.
	if _, err := os.Stat(arquivoConfig); os.IsNotExist(err) {
		fmt.Println("Arquivo de configuração (config.yml) não encontrado.")
		fmt.Print("Por favor, insira o caminho completo para o banco de dados (Ex: /home/usuário/configuracao.db): ")

		// Lê a entrada do usuário no terminal.
		leitor := bufio.NewReader(os.Stdin)
		caminho, _ := leitor.ReadString('\n')
		config.CaminhoBancoDados = strings.TrimSpace(caminho)

		// Salva a nova configuração no arquivo.
		if err := salvarConfiguracao(config); err != nil {
			return Configuracao{}, fmt.Errorf("falha ao criar config.yml: %w", err)
		}
		fmt.Printf("Arquivo de configuração salvo com o caminho: %s\n\n", config.CaminhoBancoDados)
		return config, nil
	}

	// Se o arquivo já existe, lê seu conteúdo.
	dados, err := os.ReadFile(arquivoConfig)
	if err != nil {
		return Configuracao{}, err
	}

	// Interpreta (unmarshal) o conteúdo YAML para a nossa struct 'Configuracao'.
	err = yaml.Unmarshal(dados, &config)
	if err != nil {
		return Configuracao{}, err
	}

	return config, nil
}

// ============== LÓGICA DO BANCO DE DADOS ==============

// inicializarBancoDeDados abre a conexão com o banco de dados SQLite no caminho especificado.
func inicializarBancoDeDados(caminho string) (*sql.DB, error) {
	var err error
	// Abre a conexão usando o driver 'sqlite3'.
	bancoDeDados, err = sql.Open("sqlite3", caminho)
	if err != nil {
		return nil, err
	}

	// O 'Ping' verifica se a conexão com o banco de dados é válida e está ativa.
	if err = bancoDeDados.Ping(); err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao banco de dados em '%s': %w", caminho, err)
	}

	return bancoDeDados, nil
}

// ============== LÓGICA DE CRUD - EQUIPAMENTOS ==============

// adicionarEquipamento insere um novo registro na tabela 'equipamentos'.
func adicionarEquipamento(nome, ip, cidade, tipo, vendor, devTipo string) error {
	// Prepara a instrução SQL para evitar SQL Injection.
	stmt, err := bancoDeDados.Prepare("INSERT INTO equipamentos(nome, ip, cidade, tipo, vendor, dev_tipo) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Garante que o statement será fechado ao final da função.

	// Executa a instrução preparada com os valores fornecidos.
	_, err = stmt.Exec(nome, ip, cidade, tipo, vendor, devTipo)
	return err
}

// listarEquipamentos consulta e exibe todos os registros da tabela 'equipamentos'.
func listarEquipamentos() error {
	rows, err := bancoDeDados.Query("SELECT id, nome, ip, cidade, tipo, vendor, dev_tipo FROM equipamentos")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Usa tabwriter para formatar a saída em uma tabela alinhada.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNOME\tIP\tCIDADE\tTIPO\tVENDOR\tDEV_TIPO")
	fmt.Fprintln(w, "--\t----\t--\t------\t----\t------\t--------")

	for rows.Next() {
		var id int
		// Usa sql.NullString para tratar colunas que podem ser nulas no banco.
		var nome, ip, cidade, tipo, vendor, devTipo sql.NullString
		if err := rows.Scan(&id, &nome, &ip, &cidade, &tipo, &vendor, &devTipo); err != nil {
			return err
		}
		// Imprime a linha formatada. .String converte NullString para string (vazia se for nulo).
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%s\n", id, nome.String, ip.String, cidade.String, tipo.String, vendor.String, devTipo.String)
	}
	// 'Flush' escreve o conteúdo do buffer (a tabela) no console.
	return w.Flush()
}

// deletarEquipamento remove um equipamento da tabela pelo seu ID.
func deletarEquipamento(id int) error {
	stmt, err := bancoDeDados.Prepare("DELETE FROM equipamentos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	// Verifica se alguma linha foi realmente afetada pela operação.
	linhasAfetadas, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if linhasAfetadas == 0 {
		return fmt.Errorf("nenhum equipamento encontrado com o ID %d", id)
	}
	return nil
}

// ============== LÓGICA DE CRUD - GRUPOS DE COMANDOS ==============

// adicionarGrupoComandos insere um novo registro na tabela 'grupos_comandos'.
func adicionarGrupoComandos(nome, comandos, tipoComando string) error {
	stmt, err := bancoDeDados.Prepare("INSERT INTO grupos_comandos(nome, comandos, tipo_comando) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(nome, comandos, tipoComando)
	return err
}

// listarGruposComandos consulta e exibe todos os registros da tabela 'grupos_comandos'.
func listarGruposComandos() error {
	rows, err := bancoDeDados.Query("SELECT id, nome, comandos, tipo_comando FROM grupos_comandos")
	if err != nil {
		return err
	}
	defer rows.Close()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNOME\tCOMANDOS (prévia)\tTIPO_COMANDO")
	fmt.Fprintln(w, "--\t----\t-----------------\t------------")

	for rows.Next() {
		var id int
		var nome, comandos, tipoComando sql.NullString
		if err := rows.Scan(&id, &nome, &comandos, &tipoComando); err != nil {
			return err
		}
		// Limita a exibição de comandos para não quebrar a formatação da tabela.
		comandosExibicao := comandos.String
		if len(comandosExibicao) > 50 {
			comandosExibicao = comandosExibicao[:47] + "..."
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", id, nome.String, comandosExibicao, tipoComando.String)
	}
	return w.Flush()
}

// deletarGrupoComandos remove um grupo de comandos da tabela pelo seu ID.
func deletarGrupoComandos(id int) error {
	stmt, err := bancoDeDados.Prepare("DELETE FROM grupos_comandos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	linhasAfetadas, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if linhasAfetadas == 0 {
		return fmt.Errorf("nenhum grupo de comandos encontrado com o ID %d", id)
	}
	return nil
}

// ============== DEFINIÇÃO DOS COMANDOS CLI (COBRA) ==============

// comandoRaiz é o comando principal da nossa aplicação. Os outros comandos serão "filhos" dele.
var comandoRaiz = &cobra.Command{
	Use:   "gerenciador-gcs",
	Short: "Uma CLI para gerenciar o banco de dados de configurações do GCS.",
	Long:  `Esta aplicação permite adicionar, listar e remover equipamentos e grupos de comandos da aplicação GCS. Desenvolvido por Valdemir Bezerra para o Datacenter NE.`,
}

// --- Comandos de Configuração ---
var comandoConfig = &cobra.Command{
	Use:   "config",
	Short: "Gerencia a configuração da aplicação.",
}

var comandoSetPath = &cobra.Command{
	Use:   "dir-db [novo-caminho]",
	Short: "Define um novo caminho para o banco de dados.",
	Args:  cobra.ExactArgs(1), // Exige exatamente um argumento.
	Run: func(cmd *cobra.Command, args []string) {
		config := Configuracao{CaminhoBancoDados: args[0]}
		if err := salvarConfiguracao(config); err != nil {
			log.Fatalf("Erro ao salvar nova configuração: %v", err)
		}
		fmt.Printf("Caminho do banco de dados atualizado para: %s\n", args[0])
	},
}

// --- Comandos de Equipamentos ---
var comandoEquip = &cobra.Command{
	Use:     "equip",
	Short:   "Gerencia os equipamentos.",
	Aliases: []string{"equipamento"}, // Permite usar 'equipamento' como um atalho.
}

var comandoAddEquip = &cobra.Command{
	Use:   "ad",
	Short: "Adiciona um novo equipamento.",
	Run: func(cmd *cobra.Command, args []string) {
		// Obtém os valores das flags passadas na linha de comando.
		nome, _ := cmd.Flags().GetString("nome")
		ip, _ := cmd.Flags().GetString("ip")
		cidade, _ := cmd.Flags().GetString("cidade")
		tipo, _ := cmd.Flags().GetString("tipo")
		vendor, _ := cmd.Flags().GetString("vendor")
		devTipo, _ := cmd.Flags().GetString("dev_tipo")

		if err := adicionarEquipamento(nome, ip, cidade, tipo, vendor, devTipo); err != nil {
			log.Fatalf("Erro ao adicionar equipamento: %v", err)
		}
		fmt.Println("Equipamento adicionado com sucesso!")
	},
}

var comandoListEquip = &cobra.Command{
	Use:     "list",
	Short:   "Lista todos os equipamentos.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := listarEquipamentos(); err != nil {
			log.Fatalf("Erro ao listar equipamentos: %v", err)
		}
	},
}

var comandoDeleteEquip = &cobra.Command{
	Use:     "del [ID]",
	Short:   "Deleta um equipamento pelo seu ID.",
	Aliases: []string{"rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Converte o argumento (string) para um inteiro.
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("ID inválido: '%s'. Deve ser um número.", args[0])
		}
		if err := deletarEquipamento(id); err != nil {
			log.Fatalf("Erro ao deletar equipamento: %v", err)
		}
		fmt.Printf("Equipamento com ID %d deletado com sucesso!\n", id)
	},
}

// --- Comandos de Grupos ---
var comandoGrupo = &cobra.Command{
	Use:     "grupo",
	Short:   "Gerencia os grupos de comandos.",
	Aliases: []string{"gp"},
}

var comandoAddGrupo = &cobra.Command{
	Use:   "ad",
	Short: "Adiciona um novo grupo de comandos.",
	Run: func(cmd *cobra.Command, args []string) {
		nome, _ := cmd.Flags().GetString("nome")
		comandos, _ := cmd.Flags().GetString("comandos")
		tipo, _ := cmd.Flags().GetString("tipo")

		if err := adicionarGrupoComandos(nome, comandos, tipo); err != nil {
			log.Fatalf("Erro ao adicionar grupo: %v", err)
		}
		fmt.Println("Grupo de comandos adicionado com sucesso!")
	},
}

var comandoListGrupo = &cobra.Command{
	Use:     "list",
	Short:   "Lista todos os grupos de comandos.",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := listarGruposComandos(); err != nil {
			log.Fatalf("Erro ao listar grupos: %v", err)
		}
	},
}

var comandoDeleteGrupo = &cobra.Command{
	Use:     "del [ID]",
	Short:   "Deleta um grupo de comandos pelo seu ID.",
	Aliases: []string{"rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("ID inválido: '%s'. Deve ser um número.", args[0])
		}
		if err := deletarGrupoComandos(id); err != nil {
			log.Fatalf("Erro ao deletar grupo: %v", err)
		}
		fmt.Printf("Grupo com ID %d deletado com sucesso!\n", id)
	},
}

// ============== FUNÇÃO DE INICIALIZAÇÃO E FUNÇÃO PRINCIPAL ==============

// A função init() é executada automaticamente pelo Go antes da função main().
// É o local ideal para configurar a estrutura de comandos e flags da nossa CLI.
func init() {
	// 'OnInitialize' agenda uma função para ser executada depois que as flags forem
	// processadas, mas antes que o comando principal (Run) seja executado.
	// Usamos para carregar a configuração e iniciar o banco de dados.
	cobra.OnInitialize(func() {
		cfg, err := carregarConfiguracao()
		if err != nil {
			log.Fatalf("Erro ao carregar configuração: %v", err)
		}
		if cfg.CaminhoBancoDados == "" {
			log.Fatal("O caminho do banco de dados não pode ser vazio. Use 'config set-path' para definir.")
		}
		bancoDeDados, err = inicializarBancoDeDados(cfg.CaminhoBancoDados)
		if err != nil {
			log.Fatalf("Erro ao inicializar banco de dados: %v", err)
		}
	})

	// Monta a hierarquia de comandos. Adicionamos os subcomandos ao comando raiz.
	comandoRaiz.AddCommand(comandoConfig, comandoEquip, comandoGrupo)

	// Adiciona subcomandos de 'config'.
	comandoConfig.AddCommand(comandoSetPath)

	// Adiciona subcomandos de 'equip' e define suas flags (parâmetros).
	comandoEquip.AddCommand(comandoAddEquip, comandoListEquip, comandoDeleteEquip)
	comandoAddEquip.Flags().String("nome", "", "Nome do equipamento")
	comandoAddEquip.Flags().String("ip", "", "Endereço IP do equipamento")
	comandoAddEquip.Flags().String("cidade", "", "Cidade do equipamento")
	comandoAddEquip.Flags().String("tipo", "", "Tipo do equipamento (ex: OLT, Switch)")
	comandoAddEquip.Flags().String("vendor", "", "Fabricante (ex: Huawei, Cisco)")
	comandoAddEquip.Flags().String("dev_tipo", "", "Modelo específico do equipamento")
	comandoAddEquip.MarkFlagRequired("nome") // Torna a flag --nome obrigatória.
	comandoAddEquip.MarkFlagRequired("ip")   // Torna a flag --ip obrigatória.

	// Adiciona subcomandos de 'grupo' e define suas flags.
	comandoGrupo.AddCommand(comandoAddGrupo, comandoListGrupo, comandoDeleteGrupo)
	comandoAddGrupo.Flags().String("nome", "", "Nome do grupo de comandos")
	comandoAddGrupo.Flags().String("comandos", "", "Comandos a serem executados, separados por ';'")
	comandoAddGrupo.Flags().String("tipo", "", "Tipo de comando (ex: consulta, configuracao)")
	comandoAddGrupo.MarkFlagRequired("nome")
	comandoAddGrupo.MarkFlagRequired("comandos")
}

// A função main() é o ponto de entrada de qualquer programa Go.
func main() {
	// 'defer' agenda uma chamada de função para ser executada assim que a função
	// ao redor (neste caso, main) terminar. É a forma idiomática de garantir
	// que recursos como conexões de banco de dados sejam fechados corretamente.
	if bancoDeDados != nil {
		defer bancoDeDados.Close()
	}

	// 'Execute' inicia o processamento dos argumentos da linha de comando pela biblioteca Cobra.
	if err := comandoRaiz.Execute(); err != nil {
		os.Exit(1) // Encerra o programa com um código de erro se algo falhar.
	}
}
