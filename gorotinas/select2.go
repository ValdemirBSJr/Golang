/*
E se você não quiser que sua goroutine fique bloqueada esperando?
Você pode usar o caso default. Isso é útil para "tentar" receber ou enviar de um
canal sem parar a execução.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	canalDeTarefas := make(chan string)

	// Lança uma goroutine que envia uma tarefa DEPOIS de 2 segundos.
	// Inicia junto do for
	go func() {
		time.Sleep(2 * time.Second)
		// Depois dos 2 segundos o canal recebe uma string e entra no case tarefa e finaliza
		canalDeTarefas <- "Processar relatório final"
	}()

	for {
		// Este loop vai executar muitas vezes.
		select {
		case tarefa := <-canalDeTarefas:
			fmt.Printf("Tarefa recebida: '%s'. encerrando.\n", tarefa)
			return //encerra a funcao main
		default:
			// Se NÃO houver nada para ler em `canalDeTarefas`, o caso `default`
			// é executado IMEDIATAMENTE. O `select` não bloqueia.
			fmt.Println("Nenhuma tarefa no momento, fazendo outro trabalho...")
			time.Sleep(500 * time.Millisecond) // Simula outro trabalho
		}
	}
}
