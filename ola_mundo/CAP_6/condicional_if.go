package main

import (
	"fmt"
)

func main() {

	x := 10

	if x < 100 {
		fmt.Println("Olllarr")
	}

	if !(x > 100) {
		fmt.Println("É menor que 100")
	}

	if y := 5; !(y >= 100) {
		fmt.Println("E menor ou igual a 100")
	}

	if z := 50; z > 100 {
		fmt.Println("Zê é maior que cem")
	} else if z < 10 {
		fmt.Println("Zê é menor que 10")
	} else {
		fmt.Println("Zê não é menor que deis nem maior que 100")
	}
}
