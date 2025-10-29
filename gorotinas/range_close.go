/*
close(canal) => É uma função que sinaliza que nenhum valor a mais será enviado para aquele canal.
É como o gerente de uma fábrica desligando a esteira rolante no final do dia.

Quem deve usar? Uma regra de ouro em Go:
apenas a goroutine emissora (o produtor) deve fechar um canal. Se a receptora fechar um canal
e o emissor tentar enviar dados receberá panic.

for item := range canal => O loop irá ler cada item que chega no canal e atribuí-lo à variável item.
Ele automaticamente detecta quando um canal foi fechado. Tem que ter o close senao o for bloqueara
para sempre a rotina eseprando o emissor.

*/

package main

import (
	"fmt"
	"time"
)

// `produzirTarefas` simula um trabalhador que executa uma série de tarefas
// e reporta o status através de um canal.
func produzirTarefas(listaDeTarefas []string, canalDeStatus chan<- string) {
	// A sintaxe `chan<-` reforça que esta função SÓ ENVIA para o canal.

	// `defer` garante que o canal será fechado quando a função terminar,
	// não importa o que aconteça. É a forma mais segura de garantir o fechamento.
	// Este é o sinal crucial para o `for range` saber quando parar.
	defer close(canalDeStatus)

	// Itera sobre a lista de tarefas a serem feitas.
	for _, tarefa := range listaDeTarefas {
		// Simula um tempo de processamento para cada tarefa.
		time.Sleep(1 * time.Second)

		// Cria uma mensagem de status.
		status := fmt.Sprintf("Tarefa '%s' concluída!", tarefa)

		// Envia o status para o canal. A goroutine consumidora receberá isso.
		canalDeStatus <- status
	}

	// Após o loop terminar, a função chega ao fim.
	// A instrução `defer close(canalDeStatus)` é executada agora.
	fmt.Println(">>> Produtor: Todas as tarefas foram enviadas. Fechando o canal. <<<")
}

func main() {
	// A lista de trabalho que nossa goroutine produtora irá executar.
	tarefas := []string{"Carregar Dados", "Processar Imagens", "Gerar Relatório", "Enviar Email"}

	// Criamos o canal. Poderia ser bufferizado ou não.
	// Com um buffer, o produtor poderia adiantar algumas tarefas
	// mesmo que o consumidor esteja lento.
	canalDeTarefas := make(chan string)

	// Lançamos a goroutine produtora.
	// Ela começa a trabalhar em segundo plano.
	go produzirTarefas(tarefas, canalDeTarefas)

	fmt.Println("Consumidor: Aguardando status das tarefas...")

	// O LOOP `for range`:
	// A goroutine `main` (a consumidora) vai bloquear aqui, esperando por
	// valores do `canalDeTarefas`.
	// A cada valor recebido, o corpo do loop é executado.
	// O loop só terminará quando o canal for fechado pela goroutine produtora
	// e não houver mais nenhum valor no buffer para ser lido.
	for status := range canalDeTarefas {
		fmt.Printf("Consumidor: Recebido status -> %s\n", status)
	}

	// Esta linha só será executada depois que o canal for fechado e o loop terminar.
	fmt.Println("Consumidor: O canal foi fechado. Todas as tarefas foram recebidas. Programa encerrado.")
}
