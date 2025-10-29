package main

import (
	"fmt"
)

func main() {
	// for comum
	for x := 0; x < 2; x++ {
		fmt.Println(x)
	}

	//como nao temos while, podemos fazer assim
	w := 0
	for w < 4 {
		fmt.Println("Emulando while com for")
		w++
	}

	x := 0
	for {
		if x < 10 {
			x++
			fmt.Println("X é menor que 10")
		} else {
			fmt.Println("X é maior que 10, to fora!")
			break
		}
	}

	fmt.Println("Fora do loop")

	par := 0
	for par <= 10 {
		par++
		if par%2 != 0 {
			continue
			//avança apenas essa interação pra o proximo passo
		}
		fmt.Println(par)
	}

}
