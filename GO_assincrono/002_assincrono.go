package main

import (
	"encoding/json" // Pacote para processar JSON
	"fmt"
	"io"       // Pacote para ler a resposta da requisição
	"net/http" // Pacote para fazer as chamadas HTTP
)

// --- Estruturas para Processar o JSON de cada API ---

// Estrutura para extrair apenas a informação que nos interessa da API de clima (wttr.in)
// Os campos da struct precisam ser exportados (começar com letra maiúscula)
type RespostaClima struct {
	CondicaoAtual []struct {
		TemperaturaCelsius string `json:"temp_C"`
		SensacaoTermica    string `json:"FeelsLikeC"`
	} `json:"current_condition"`
}

// Estrutura para extrair a cotação do Bitcoin da API da CoinGecko
type RespostaBitcoin struct {
	Bitcoin struct {
		PrecoBRL float64 `json:"brl"`
	} `json:"bitcoin"`
}

// --- Funções Concorrentes ---

// buscarClima faz uma chamada real à API wttr.in para obter o clima de uma cidade.
func buscarClima(cidade string, canal chan string) {
	// Monta a URL da API. O formato `?format=j1` pede uma resposta em JSON.
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", cidade)

	// Faz a requisição HTTP GET.
	resposta, err := http.Get(url)
	if err != nil {
		// Se houver um erro de rede, envia uma mensagem de erro pelo canal.
		canal <- fmt.Sprintf("Erro ao buscar clima: %v", err)
		return
	}
	// `defer` garante que o corpo da resposta será fechado ao final da função,
	// liberando a conexão de rede. É uma prática essencial em Go.
	defer resposta.Body.Close()

	// Lê todo o corpo da resposta.
	corpo, err := io.ReadAll(resposta.Body)
	if err != nil {
		canal <- fmt.Sprintf("Erro ao ler resposta do clima: %v", err)
		return
	}

	// Agora, processamos o JSON.
	var dadosClima RespostaClima
	// `json.Unmarshal` "desempacota" o JSON (em bytes) para dentro da nossa struct.
	err = json.Unmarshal(corpo, &dadosClima)
	if err != nil {
		canal <- fmt.Sprintf("Erro ao processar JSON do clima: %v", err)
		return
	}

	// Pega a primeira condição atual (geralmente a única).
	if len(dadosClima.CondicaoAtual) > 0 {
		condicao := dadosClima.CondicaoAtual[0]
		resultadoFormatado := fmt.Sprintf("Clima em %s: %s°C (Sensação de %s°C)",
			cidade, condicao.TemperaturaCelsius, condicao.SensacaoTermica)
		canal <- resultadoFormatado
	} else {
		canal <- "Não foi possível obter a condição do tempo."
	}
}

// buscarCotacaoBitcoin faz uma chamada real à API da CoinGecko.
func buscarCotacaoBitcoin(canal chan string) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=brl"

	resposta, err := http.Get(url)
	if err != nil {
		canal <- fmt.Sprintf("Erro ao buscar cotação do Bitcoin: %v", err)
		return
	}
	defer resposta.Body.Close()

	corpo, err := io.ReadAll(resposta.Body)
	if err != nil {
		canal <- fmt.Sprintf("Erro ao ler resposta do Bitcoin: %v", err)
		return
	}

	var dadosBitcoin RespostaBitcoin
	err = json.Unmarshal(corpo, &dadosBitcoin)
	if err != nil {
		canal <- fmt.Sprintf("Erro ao processar JSON do Bitcoin: %v", err)
		return
	}

	// Formata o resultado para ter uma aparência de moeda. `%.2f` limita a 2 casas decimais.
	resultadoFormatado := fmt.Sprintf("Cotação Bitcoin: R$ %.2f", dadosBitcoin.Bitcoin.PrecoBRL)
	canal <- resultadoFormatado
}

func main() {
	fmt.Println("Iniciando buscas concorrentes nas APIs...")

	// 1. Cria o canal que receberá os resultados.
	canalDeResultados := make(chan string)

	// 2. Dispara as duas buscas em Goroutines.
	// O programa não espera, ele lança as duas e continua.
	go buscarClima("Joao+Pessoa", canalDeResultados)
	go buscarCotacaoBitcoin(canalDeResultados)

	// 3. Aguarda e recebe os resultados do canal.
	// O programa vai pausar aqui até o primeiro resultado chegar,
	// seja ele qual for (clima ou bitcoin).
	resultado1 := <-canalDeResultados
	fmt.Println("Recebido:", resultado1)

	// Em seguida, ele vai pausar novamente esperando o segundo resultado.
	resultado2 := <-canalDeResultados
	fmt.Println("Recebido:", resultado2)

	fmt.Println("\nAmbas as buscas foram concluídas!")
}
