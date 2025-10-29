/*
Usando Canais como "Pipeline"
Este é o padrão mais poderoso e comum, ideal quando a tarefaDois
depende do resultado (do output) da tarefaUm para poder fazer o seu próprio trabalho.

Criamos uma esteira (pipeline), onde a saída de uma goroutine se torna a entrada da próxima.

SAÍDA GARANTIDA:
Pipeline iniciado. Aguardando a conclusão da tarefa consumidora.
Tarefa 1 (Geradora): Iniciando processamento...
Tarefa 2 (Consumidora): Pronta para receber dados...
  -> Tarefa 2 (Consumidora): Recebido e finalizando -> Item 'A' processado
  -> Tarefa 2 (Consumidora): Recebido e finalizando -> Item 'B' processado
  -> Tarefa 2 (Consumidora): Recebido e finalizando -> Item 'C' processado
Tarefa 1 (Geradora): Todos os itens foram processados e enviados.
Tarefa 2 (Consumidora): O canal foi fechado. Finalizando.
Pipeline finalizado.

Este padrão garante não apenas a ordem, mas também permite o processamento
em fluxo (streaming), onde a `tarefaDois` começa a trabalhar assim que o
primeiro item da `tarefaUm` fica pronto, sem precisar esperar que todos
os itens sejam processados primeiro.

RESUMO:
Para ordenar tarefas dependentes (B precisa do resultado de A):
Use canais para criar um pipeline. A primeira tarefa (produtora) envia seus resultados
para um canal e o fecha no final. A segunda tarefa (consumidora) lê desse canal usando um
loop for range até que ele seja fechado.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// tarefaUm_Geradora produz dados e os envia para um canal de saída.
// Por convenção, a goroutine que produz os dados é responsável por fechar o canal.
func tarefaUm_Geradora(dadosParaProcessar []string, saida chan<- string) {
	// A sintaxe `chan<-` indica que esta função só envia para o canal `saida`.
	defer close(saida) // Garante que o canal será fechado ao final da função.

	fmt.Println("Tarefa 1 (Geradora): Iniciando processamento...")
	for _, item := range dadosParaProcessar {
		// Simula um trabalho pesado para cada item.
		time.Sleep(500 * time.Millisecond)
		resultado := fmt.Sprintf("Item '%s' processado", item)
		// Envia o resultado para o canal de saída.
		saida <- resultado
	}
	fmt.Println("Tarefa 1 (Geradora): Todos os itens foram processados e enviados.")
}

// tarefaDois_Consumidora lê dados de um canal de entrada e os processa.
func tarefaDois_Consumidora(entrada <-chan string, wg *sync.WaitGroup) {
	// A sintaxe `<-chan` indica que esta função só recebe do canal `entrada`.
	defer wg.Done()

	fmt.Println("Tarefa 2 (Consumidora): Pronta para receber dados...")
	// A FORMA IDIOMÁTICA DE LER UM CANAL ATÉ O FIM: `for range`.
	// Este loop irá ler do canal `entrada` até que ele seja fechado
	// pela goroutine produtora (`tarefaUm_Geradora`).
	for resultadoDaTarefaUm := range entrada {
		fmt.Printf("  -> Tarefa 2 (Consumidora): Recebido e finalizando -> %s\n", resultadoDaTarefaUm)
		// Aqui poderíamos fazer um trabalho adicional com o resultado.
	}
	fmt.Println("Tarefa 2 (Consumidora): O canal foi fechado. Finalizando.")
}

func main() {
	// Os dados iniciais que a primeira tarefa irá processar.
	dadosIniciais := []string{"A", "B", "C"}

	// Criamos o canal que conectará a Tarefa 1 com a Tarefa 2.
	canalDoPipeline := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1) // `main` só precisa esperar a última etapa do pipeline terminar.

	// Lançamos as duas goroutines, conectando-as através do canal.
	go tarefaUm_Geradora(dadosIniciais, canalDoPipeline)
	go tarefaDois_Consumidora(canalDoPipeline, &wg)

	fmt.Println("Pipeline iniciado. Aguardando a conclusão da tarefa consumidora.")
	wg.Wait()
	fmt.Println("Pipeline finalizado.")
}
