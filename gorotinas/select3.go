/*
Um dos usos mais comuns do select é para evitar que uma
goroutine espere para sempre por uma operação. Combinamos um case para a
operação desejada com um case para um temporizador.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	canalDaOperacao := make(chan string)

	// Lança uma goroutine que simula uma operação demorada (3 segundos).
	go func() {
		time.Sleep(3 * time.Second)
		canalDaOperacao <- "Operação concluída com sucesso"
	}()

	fmt.Println("Aguardando pela operação com um timeout de 2 segundos...")

	select {
	// Caso 1: A operação terminou a tempo.
	case resultado := <-canalDaOperacao:
		fmt.Printf("Resultado: %s\n", resultado)
	// Caso 2: O tempo se esgotou.
	// `time.After(duration)` retorna um canal que recebe um valor
	// após a duração especificada.
	case <-time.After(2 * time.Second):
		fmt.Println("A operação excedeu o tempo limite (timeout)!")
	}

}
