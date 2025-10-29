package main

import (
	"fmt"
)

func main() {
	x := 2
	if x == 2 || x == 3 {
		fmt.Println("X é dois ou três")
	}

	y := 6
	if y%2 == 0 && y%3 == 0 {
		fmt.Println("Y é múltiplo de 2 e também de 3")
	}

	z := 2
	if !(z == 2) || z < 3 {
		fmt.Println("Z não é igual a 2 ou menor que 3")
	}

}
