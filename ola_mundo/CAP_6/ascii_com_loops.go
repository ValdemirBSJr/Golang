package main

import (
	"fmt"
)

func main() {
	numeros := 33

	for numeros <= 122 {
		fmt.Printf("O número impresso é %d\t hexa do dele: %#x\t unicode dele: %#U\n", numeros, numeros, numeros)
		numeros++
	}
}
