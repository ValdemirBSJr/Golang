/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
  - "Nome", "Sobrenome", "Hobby favorito"

- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import (
	"fmt"
)

func main() {

	slice_multimodal := [][]string{

		[]string{
			"Toin",
			"Das crui",
			"Babão de político",
		},

		[]string{
			"Arnaldinn",
			"Melado de Nego",
			"Galã de cabaré rasga zorba",
		},

		[]string{
			"Vadalco",
			"Ribonbo",
			"Polidor de chifre de corno",
		},
	}

	for _, v := range slice_multimodal {
		fmt.Println(v)
	}

	fmt.Printf("\n\n")

	for _, v := range slice_multimodal {
		fmt.Println(v[0])
		for _, item := range v {
			if item != v[0] {
				fmt.Println("\t", item)
			}

		}
	}

}
