/*
Go utiliza dois conceitos principais que simplificam enormemente o processo:
Goroutines: São como threads, mas muito mais leves. É possível ter milhares ou
até milhões delas rodando simultaneamente sem sobrecarregar o sistema. Iniciar
uma é tão simples quanto usar a palavra-chave go na frente de uma chamada de função.

Channels (Canais): São "tubos" tipados através dos quais as goroutines podem se comunicar e sincronizar.
Em vez de compartilhar memória para se comunicar (o que exige locks), as goroutines enviam e recebem valores
através dos canais. A filosofia de Go é: "Não comunique compartilhando memória; em vez disso, compartilhe memória comunicando."
*/
package main

import (
	"fmt"
	"time" // Pacote para simular uma demora
)

// Esta função simula a busca de dados em uma fonte externa.
// Ela recebe um nome de fonte e um canal para enviar a resposta.
func buscarDados(fonte string, canal chan string) {
	fmt.Printf("Buscando dados na fonte: %s...\n", fonte)

	// Simula um tempo de espera, como uma chamada de rede.
	// A fonte "API Principal" demora 2 segundos.
	if fonte == "API Principal" {
		time.Sleep(time.Second * 2)
	} else {
		// A fonte "API Secundária" demora 1 segundo.
		time.Sleep(time.Second * 1)
	}

	// Após a "busca" terminar, enviamos uma mensagem de resultado
	// de volta pelo canal. A seta `<-` indica o envio de um valor para o canal.
	canal <- "Dados recebidos de: " + fonte
}

func main() {
	fmt.Println("Iniciando a busca por dados...")

	// Criamos um canal. Este canal servirá como ponto de encontro
	// para as nossas goroutines. Ele transportará valores do tipo `string`.
	canalDeResultados := make(chan string)

	// --- A MÁGICA DA CONCORRÊNCIA ACONTECE AQUI ---

	// 1. Iniciamos a primeira busca em uma Goroutine.
	// A palavra `go` na frente da chamada da função faz com que ela execute
	// de forma concorrente, em "background". O programa principal NÃO espera
	// ela terminar e continua para a próxima linha imediatamente.
	go buscarDados("API Principal", canalDeResultados)

	// 2. Iniciamos a segunda busca em outra Goroutine.
	// Esta também começa a executar imediatamente, em paralelo com a primeira.
	go buscarDados("API Secundária", canalDeResultados)

	// --- SINCRONIZAÇÃO ---

	// O programa principal agora precisa esperar pelas respostas.
	// Ele vai "ouvir" o canal. A operação de receber `<-canalDeResultados`
	// é bloqueante: ela pausa a execução da função `main` até que um valor
	// seja enviado para o canal.
	fmt.Println("Aguardando o primeiro resultado...")
	resultado1 := <-canalDeResultados
	fmt.Println(resultado1)

	fmt.Println("Aguardando o segundo resultado...")
	resultado2 := <-canalDeResultados
	fmt.Println(resultado2)

	fmt.Println("Buscas finalizadas. O programa pode continuar.")
}
