//slice e mais flexivel que array. eu consigo mudar o tamanho dele. array uma vez setado, nao

package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	slice_str := []string{"banana", "maçã", "jaca"}

	for índice, valor := range slice_str {
		fmt.Println("No índice ", índice, ". Temos o valor: ", valor)
	}

	//slice_str[3] = "melancia" //aqui vai dar erro
	slice_str = append(slice_str, "melancia") //Aqui nao da erro

	for índice, valor := range slice_str {
		fmt.Println("No índice ", índice, ". Temos o valor: ", valor)
	}

	for _, valor := range slice_str {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

	total := ""
	for _, valor := range slice_str {
		total += valor
		fmt.Printf("O valor cumulativo: %v\n", total)
	}

	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	fmt.Println("O valor total é: ", soma)

}
