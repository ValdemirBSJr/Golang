// GO não tem while
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	for x := 1; x <= 10; x++ {
		fmt.Print(" ", x)
	}

	fmt.Println()

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora: ", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}

		fmt.Println()
	}

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("MÊS: ", mes)

		for dia := 1; dia <= 30; dia++ {
			fmt.Print("Dia: ", dia, ", ")
		}

		fmt.Println()
	}

}
