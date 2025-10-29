package main

import (
	"fmt"
)

func main() {
	primeiro_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiro_slice)
	segundoslice := append(primeiro_slice[:2], primeiro_slice[4:]...)
	fmt.Println(segundoslice)
	fmt.Println(primeiro_slice)

	//quando vc faz o append ele faz um fatiamento em cima do slice original e altera
	// vc pode burlar isso copiando item a item com for ou
}
