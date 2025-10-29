package main

import (
	"fmt"
)

var y = 15

func main() {
	fmt.Println("Função principal")

	fmt.Println(y)

	y := 10
	qualquercoisa(y)
}

func qualquercoisa(x int) {
	fmt.Println("Função qualquer coisa:")
	fmt.Println(x)
}
