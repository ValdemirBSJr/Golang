// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)

		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

}
