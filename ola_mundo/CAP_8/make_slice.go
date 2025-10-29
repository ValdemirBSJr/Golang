package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)
	// o make serve para criar capacidades precisas
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	fmt.Println(slice, len(slice), cap(slice)) //cap = capacidade

	//adicionar 1
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice)) //cap = capacidade
}
