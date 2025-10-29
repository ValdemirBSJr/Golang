/*
Veremos um exemplo de como garantir que a `main` espere.
A forma correta de se esperar as goroutines terminarem é sync.WaitGroup.
Pense nele como um contador de tarefas ativas.

FLUXO DE EXECUÇÃO DESTE CÓDIGO IDIOMÁTICO:

 1. `main` cria um `WaitGroup`.
 2. `main` diz: "WaitGroup, espere por 1 tarefa" (`wg.Add(1)`).
 3. `main` lança `go fale("mundo", &wg)`. A goroutine "mundo" começa a executar.
 4. `main` executa `faleSemWaitGroup("Olá, ")`. Durante esse tempo, "Olá, " e "mundo"
    são impressos de forma intercalada.
 5. `faleSemWaitGroup("Olá, ")` termina.
 6. `main` chega em `wg.Wait()` e para.
 7. A goroutine "mundo" continua executando. Quando seu loop de 5 iterações termina,
    a instrução `defer wg.Done()` é executada. Isso diminui o contador do WaitGroup para 0.
 8. Como o contador é 0, `wg.Wait()` em `main` é desbloqueado.
 9. `main` imprime a mensagem final e o programa encerra de forma limpa.
*/
package main

import (
	"fmt"
	"sync" // Importamos o pacote `sync` para usar o WaitGroup
	"time"
)

// A função `fale` agora aceita um ponteiro para um WaitGroup.
// Ela precisa dele para sinalizar quando terminou seu trabalho.
func fale(s string, wg *sync.WaitGroup) {
	// `defer` é uma palavra-chave do Go que adia a execução de uma função
	// até o final da função atual.
	// Aqui, garantimos que `wg.Done()` será chamado antes de `fale` retornar,
	// não importa como ela termine. É uma prática de segurança muito comum.
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

// Função de exemplo para a main executar sem ser uma goroutine controlada.
func faleSemWaitGroup(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 1. Criamos uma variável do tipo WaitGroup.
	var wg sync.WaitGroup

	// 2. Avisamos ao WaitGroup que vamos adicionar UMA goroutine
	//    para ele esperar. O contador interno agora é 1.
	wg.Add(1)

	// Lançamos nossa goroutine, passando a ela uma referência
	// ao nosso WaitGroup (`&wg`).
	//"Ó a variavel que voce vai pegar o ponteiro é essa"
	go fale("mundo!", &wg)

	// A goroutine principal continua seu trabalho...
	// (Neste exemplo, não passamos `fale("Olá, ")` para uma goroutine,
	// mas poderíamos, bastando adicionar outro `wg.Add(1)` e `go fale(...)`).
	faleSemWaitGroup("Olá, ")

	// 3. O PONTO CHAVE: wg.Wait()
	// Esta linha BLOQUEIA a execução da goroutine `main`.
	// O programa vai parar aqui e só continuará quando o contador
	// interno do WaitGroup chegar a zero.
	// O contador só chega a zero quando a goroutine `fale("mundo!")`
	// chama `wg.Done()`.
	fmt.Println("Aguardando a goroutine 'mundo' terminar...")
	wg.Wait()

	fmt.Println("Goroutine 'mundo' terminou. Encerrando o programa.")

}
