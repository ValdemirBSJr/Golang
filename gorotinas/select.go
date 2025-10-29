/*
 Imagine que você tem dois serviços web que te enviam notificações:
 um é muito rápido e envia notícias urgentes, e o outro é mais lento e envia um resumo diário.
 Sua goroutine principal precisa ouvir ambos e processar as notícias assim que chegarem.

 Pense no select como uma instrução switch, mas para operações de canais.
 Eu tenho vários canais (canal A, canal B, canal C) que estou monitorando.
 Vou esperar aqui até que pelo menos um deles esteja pronto para uma operação
 (seja um envio ou um recebimento). O primeiro que ficar pronto, eu executo o bloco de
 código correspondente a ele.

 O select bloqueia a execução até que um dos seus case (casos) possa ser executado.

Se múltiplos case estiverem prontos ao mesmo tempo, o select escolhe um deles de
forma pseudo-aleatória para executar. Isso garante que nenhum canal seja deixado de lado
(evita starvation).

Existe um caso opcional, default, que é executado se nenhum dos outros casos estiver pronto,
tornando o select não-bloqueante.
*/

package main

import (
	"fmt"
	"time"
)

// `servicoDeNoticiasUrgentes` envia uma notícia a cada 1 segundo.
func servicoDeNoticiasUrgentes(canal chan<- string) {
	for i := 1; ; i++ {
		mensagem := fmt.Sprintf("Notícia urgente Nº %d", i)
		canal <- mensagem
		time.Sleep(1 * time.Second)
	}
}

// `servicoDeResumoDiario` envia um resumo a cada 3 segundos.
func servicoDeResumoDiario(canal chan<- string) {
	for i := 1; ; i++ {
		mensagem := fmt.Sprintf("Resumo diário Nº %d", i)
		canal <- mensagem
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Criamos um canal para cada serviço de notícias.
	canalUrgente := make(chan string)
	canalResumo := make(chan string)

	// Lançamos os dois serviços em goroutines separadas.
	go servicoDeNoticiasUrgentes(canalUrgente)
	go servicoDeResumoDiario(canalResumo)

	fmt.Println("Central de Notícias: Aguardando mensagens de todos os serviços...")

	// Loop infinito para que a `main` continue escutando.
	for {
		// O `select` vai bloquear aqui até que uma mensagem chegue
		// em `canalUrgente` OU em `canalResumo`.
		select {
		// Caso 1: Se uma mensagem puder ser lida de `canalUrgente`...
		case mensagemUrgente := <-canalUrgente:
			fmt.Printf("[URGENTE] %s\n", mensagemUrgente)

		// Caso 2: Se uma mensagem puder ser lida de `canalResumo`...
		case mensagemResumo := <-canalResumo:
			fmt.Printf("[RESUMO DIÁRIO] %s\n", mensagemResumo)

		}
	}
}
