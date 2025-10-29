//Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import (
	"fmt"
)

func main() {
	tamanho_do_cansaco := 1

	switch {
	case tamanho_do_cansaco == 0:
		fmt.Println("Que malandragem")
	case tamanho_do_cansaco == 1:
		fmt.Println("Uma gelada ia bem")
	case tamanho_do_cansaco > 1:
		fmt.Println("Ih, já era")
	}
}
