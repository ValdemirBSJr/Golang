//Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import (
	"fmt"
)

func main() {
	tamanho_do_cansaco := 2
	if tamanho_do_cansaco == 0 {
		fmt.Println("Que malandragem")
	} else if tamanho_do_cansaco == 1 {
		fmt.Println("uma gelada ia bem")
	} else {
		fmt.Println("Ih, já era...")
	}
}
