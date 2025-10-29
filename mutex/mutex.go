/*
O Problema: Condição de Corrida (Race Condition)

magine que um map é como um quadro branco. Se duas pessoas (goroutines)
tentam ler e escrever no mesmo lugar do quadro ao mesmo tempo, o resultado será uma bagunça
ilegível. Em Go, tentar escrever em um map a partir de múltiplas goroutines simultaneamente,
sem proteção, causa um erro fatal chamado panic.

O sync.Mutex (Exclusão Mútua) age como um "cadeado". A regra é simples:
só a goroutine que está com a chave (que travou o cadeado) pode mexer no quadro
branco (os dados). As outras precisam esperar na fila pela sua vez.
*/
package main

import (
	"fmt"
	"sync"
)

// ContadorSeguro é um contador que pode ser usado de forma segura por múltiplas goroutines.
type ContadorSeguro struct {
	// mu é um Mutex, que significa 'exclusão mútua'.
	// Ele garante que apenas uma goroutine possa acessar o map 'valores' de cada vez.
	mu      sync.Mutex
	valores map[string]int
}

// Incrementar incrementa o contador para uma dada chave de forma segura.
func (c *ContadorSeguro) Incrementar(chave string) {
	// Trava o mutex. A partir daqui, nenhuma outra goroutine pode executar
	// este trecho de código até que chamemos Unlock().
	// Se outra goroutine tentar chamar Lock(), ela ficará bloqueada esperando.
	c.mu.Lock()

	// O código entre Lock() e Unlock() é chamado de "seção crítica".
	// Apenas uma goroutine por vez pode estar aqui.
	c.valores[chave]++

	// Libera o mutex, permitindo que a próxima goroutine em espera possa
	// adquirir o lock e executar a seção crítica.
	c.mu.Unlock()
}

// Valor retorna o valor atual do contador para uma dada chave.
func (c *ContadorSeguro) Valor(chave string) int {
	// Trava o mutex para garantir que estamos lendo um valor consistente,
	// sem que outra goroutine o modifique no meio da leitura.
	c.mu.Lock()

	// 'defer' agenda a execução de c.mu.Unlock() para o exato momento em que
	// a função 'Valor' estiver prestes a retornar.
	// Esta é a forma preferida e mais segura de liberar um mutex, pois garante
	// que ele SEMPRE será liberado, mesmo se a função tiver múltiplos retornos
	// ou sofrer um 'panic'.
	defer c.mu.Unlock()

	return c.valores[chave]

}

func main() {
	// Cria uma instância do nosso contador seguro.
	contador := ContadorSeguro{valores: make(map[string]int)}

	// Um WaitGroup é usado para esperar que um conjunto de goroutines termine.
	// Pense nele como um contador de goroutines ativas.
	var wg sync.WaitGroup

	//Vamos lançar 1000 goroutines
	for i := 0; i < 1000; i++ {
		// Avisa ao WaitGroup que estamos adicionando 1 goroutine ao contador.
		wg.Add(1)

		go func() {
			// 'defer wg.Done()' garante que avisaremos ao WaitGroup que esta
			// goroutine terminou, não importa como ela termine.
			// Done() decrementa o contador do WaitGroup em 1.
			defer wg.Done()

			// Cada goroutine incrementa a mesma chave.
			contador.Incrementar("minhaChave")

		}()

		// wg.Wait() bloqueia a execução da função 'main' até que o contador
		// interno do WaitGroup chegue a zero. Ou seja, espera todas as
		// goroutines que chamaram wg.Add(1) também chamarem wg.Done().
		wg.Wait()

		// Neste ponto, temos a garantia de que TODAS as 1000 goroutines
		// terminaram suas execuções.
		// Agora é seguro ler o valor final.
		fmt.Println("Valor final:", contador.Valor("minhaChave"))

	} //for
} //main
