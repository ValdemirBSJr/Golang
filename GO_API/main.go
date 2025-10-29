package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Estrutura para a nossa resposta da API
type MensagemAPI struct {
	Texto     string `json:"texto"`
	Timestamp string `json:"timestamp"`
	Frase     string `json:"frase"`
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

func main() {
	// Semente para o gerador de números aleatórios. Essencial para aleatoriedade real.
	rand.Seed(time.Now().UnixNano())

	// Registra nosso único handler para a rota principal da API
	http.HandleFunc("/api/mensagem", mensagemHandler)

	fmt.Println("🚀 Servidor Go (versão final) rodando na porta 8080...")
	fmt.Println("Endpoint disponível: http://localhost:8080/api/mensagem")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
