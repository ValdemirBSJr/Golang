package main

import (
	"fmt"
)

func main() {

	basica()
	argumento("tarde")

	valor := soma(10, 1)

	fmt.Println(valor)

	total, quantidade := variatica(10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantidade)
}

// nome da func, tipo dos parametros, tipo do retorno
func soma(x, y int) int {
	return x + y
}

func basica() {
	fmt.Println("Bom dia")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi. bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde.")
	} else {
		fmt.Println("Oi, ba noite.")
	}
}

// funcao variatica, como kwargs, retorna 2 inteiros, como tem 2 retornos vai entre parentesis
// variatica(s string, x ... int)(int, int, string)
// o parametro variatico tem que ser sempre o ultimo
func variatica(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
