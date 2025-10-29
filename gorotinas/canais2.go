/*
Este padrão é ideal quando a tarefaDois não precisa de um resultado da tarefaUm,
ela apenas precisa saber que a tarefaUm terminou antes que ela possa começar.

Usamos um canal para enviar um sinal de finalização. A forma mais eficiente e idiomática
para um canal que só sinaliza é chan struct{} (um canal de structs vazias), pois não aloca
memória alguma.

SAÍDA GARANTIDA:
Iniciando tarefa 1 (demora 2 segundos)...
Tarefa 2 está pronta, aguardando sinal da Tarefa 1...
Goroutines lançadas. Aguardando a conclusão de todas as tarefas.
...Tarefa 1 finalizada!
Sinal recebido! Iniciando tarefa 2 (demora 1 segundo)...
...Tarefa 2 finalizada!
Todas as tarefas foram concluídas.

Note que a `tarefaDois` SEMPRE inicia seu trabalho APÓS a `tarefaUm` ter finalizado.
A sincronização foi orquestrada com sucesso pelo canal.

RESUMO:
Para ordenar tarefas independentes (A termina -> B começa):
Use um canal (chan struct{}) como um sinal. A primeira tarefa envia no canal ao terminar,
e a segunda bloqueia esperando para receber.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// tarefaUm faz algum trabalho demorado...
func tarefaUm(sinal chan<- struct{}) {
	// A sintaxe `chan<-` significa que esta função só pode ENVIAR para o canal.
	fmt.Println("Iniciando tarefa 1 (demora 2 segundos)...")
	time.Sleep(2 * time.Second)
	fmt.Println("...Tarefa 1 finalizada!")

	// Envia um sinal para o canal para indicar que terminou.
	// Uma struct vazia `struct{}{} ` é enviada. O valor não importa, apenas o ato de enviar.
	// Outra forma idiomática de sinalizar é usando `close(sinal)`. Veremos isso depois.
	sinal <- struct{}{}
}

// tarefaDois espera o sinal da tarefaUm para começar.
func tarefaDois(sinal <-chan struct{}) {
	// A sintaxe `<-chan` significa que esta função só pode RECEBER do canal.

	fmt.Println("Tarefa 2 está pronta, aguardando sinal da Tarefa 1...")
	// ESTA LINHA É BLOQUEANTE.
	// A execução para aqui até que um valor seja enviado para o canal `sinal`.
	<-sinal
	fmt.Println("Sinal recebido! Iniciando tarefa 2 (demora 1 segundo)...")
	time.Sleep(1 * time.Second)
	fmt.Println("...Tarefa 2 finalizada!")
}

func main() {
	// Cria o canal que servirá como ponte de sinalização entre as goroutines.
	canalDeSinal := make(chan struct{})

	// Usamos um WaitGroup para garantir que a `main` espere por todas as tarefas.
	var wg sync.WaitGroup
	wg.Add(2) // Vamos esperar por 2 goroutines.

	// Lançamos as duas goroutines.
	// Elas começam a ser agendadas pelo Go Scheduler ao mesmo tempo.
	go func() {
		defer wg.Done()
		tarefaUm(canalDeSinal)
	}()

	go func() {
		defer wg.Done()
		tarefaDois(canalDeSinal)
	}()

	fmt.Println("Goroutines lançadas. Aguardando a conclusão de todas as tarefas.")
	// `main` espera aqui até que ambas as goroutines chamem `wg.Done()`.
	wg.Wait()
	fmt.Println("Todas as tarefas foram concluídas.")
}
