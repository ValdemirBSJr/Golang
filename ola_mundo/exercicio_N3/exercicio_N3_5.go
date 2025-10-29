//Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import (
	"fmt"
)

func main() {
	var x int = 10

	for x <= 100 {

		fmt.Println(x % 4)
		x++

	}
}
