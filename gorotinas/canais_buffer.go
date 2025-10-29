/*
Um canal sem buffer (make(chan int)) é uma esteira minúscula
onde uma pessoa entrega um item diretamente na mão da outra. Se o
receptor não estiver lá com a mão estendida, o emissor para e espera.

Um canal com buffer (make(chan int, 2))
é uma esteira rolante que tem espaço para alguns itens.
O emissor pode colocar itens na esteira e sair, sem precisar esperar pelo receptor,
contanto que haja espaço livre na esteira.

As regras de bloqueio mudam:

Envio: Uma goroutine que envia para um canal bufferizado só bloqueia se o buffer estiver cheio.
Se houver espaço, ela coloca o item no buffer e continua sua execução imediatamente.

Recebimento: Uma goroutine que recebe de um canal bufferizado só bloqueia se o buffer estiver vazio.
Se houver itens, ela pega um e continua sua execução imediatamente.


Quando Usar Canais Bufferizados?
Absorver Picos de Trabalho (Bursts)
Evitar Impasses (Deadlocks) Simples
Aumentar a Taxa de Transferência (Throughput)

 Cuidado: O uso de canais bufferizados torna a sincronização menos explícita.
 É uma troca: você ganha em desacoplamento e performance, mas perde a garantia de que,
 quando um envio é feito, o receptor já está ciente dele. Use-os quando você entende essa
 troca e precisa do comportamento de fila que eles oferecem.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// A SINTAXE:
	// Criamos um canal de strings com um buffer de capacidade 2.
	// Isso significa que podemos enviar até 2 valores para este canal
	// ANTES que o emissor precise bloquear.
	canal := make(chan string, 2)

	// --- ENVIANDO DADOS ---

	// Enviamos o primeiro valor.
	// Como o buffer tem espaço (capacidade 2, uso 0), esta operação NÃO bloqueia.
	// A goroutine `main` continua imediatamente.
	canal <- "olá"
	fmt.Println("Enviado 'olá' para o canal sem bloquear.")

	// Enviamos o segundo valor.
	// O buffer ainda tem espaço (capacidade 2, uso 1), então esta operação
	// também NÃO bloqueia.
	canal <- "mundo"
	fmt.Println("Enviado 'mundo' para o canal sem bloquear. Buffer agora está cheio.")

	// AGORA O BUFFER ESTÁ CHEIO (capacidade 2, uso 2).
	// A próxima tentativa de envio irá bloquear.
	fmt.Println("Próximo envio irá bloquear até que um item seja lido...")

	// Para demonstrar o bloqueio sem parar o programa,
	// faremos o próximo envio em uma goroutine separada.
	go func() {
		fmt.Println("Goroutine extra: tentando enviar '!'...")
		canal <- "!" // <-- ESTA LINHA VAI BLOQUEAR A GOROUTINE EXTRA
		fmt.Println("Goroutine extra: '!' finalmente foi enviado!")
	}()

	// Damos um pequeno tempo para a goroutine extra tentar enviar e bloquear.
	time.Sleep(3 * time.Second)
	fmt.Println("---------------------------------")

	// --- RECEBENDO DADOS ---

	// A goroutine `main` agora vai ler o primeiro valor do canal.
	// Como o buffer não está vazio (contém "olá" e "mundo"),
	// esta operação NÃO bloqueia.
	primeiraMensagem := <-canal
	fmt.Printf("Recebido '%s' do canal.\n", primeiraMensagem)
	fmt.Println("Espaço foi liberado no buffer.")

	// Damos mais um tempo para observar o que acontece.
	// Como `main` leu um valor, o buffer agora tem espaço livre.
	// A goroutine extra que estava bloqueada tentando enviar "!"
	// será desbloqueada e conseguirá completar seu envio.
	time.Sleep(3 * time.Second)

	// Agora, vamos ler os valores restantes.
	segundaMensagem := <-canal
	terceiraMensagem := <-canal
	fmt.Printf("Recebido '%s' do canal.\n", segundaMensagem)
	fmt.Printf("Recebido '%s' do canal.\n", terceiraMensagem)

}
