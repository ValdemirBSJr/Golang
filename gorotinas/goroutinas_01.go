/*
Pense em uma goroutine como uma "tarefa" que seu programa pode executar.
A grande vantagem é que o Go pode executar milhares, ou até milhões,
dessas tarefas ao mesmo tempo (de forma concorrente) sem o peso e a complexidade
das threads tradicionais que você encontra em outras linguagens.

A "mágica" acontece com a palavra-chave go. Quando você a coloca na frente de
uma chamada de função, você está dizendo ao Go:

"Execute esta função, mas não espere ela terminar. Crie uma nova tarefa
(goroutine) para ela e continue executando o resto do meu código imediatamente."

IMPORTANTE:
Um programa Go encerra quando a função `main` (a goroutine principal) termina.
Ele NÃO espera outras goroutines terminarem.
Por isso a saída será:
mundo!
Olá,
Olá,
O go não espera os milissegundos das repetições
*/
package main

import (
	"fmt"
	"time"
)

func fale(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go fale("mundo!")
	fale("Olá, ")
}
