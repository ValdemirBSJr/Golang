/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
  - Nome
  - Sobrenome
  - Sabores favoritos de sorvete

- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import (
	"fmt"
)

type pessoa struct {
	nome               string
	sobrenome          string
	sorvetes_favoritos [3]string //pode ser assim tbm []string e fica em aberto o tamanho
}

func main() {

	pessoa1 := pessoa{"Roberto", "Da costa", [3]string{"balmilha", "chuculate", "murango"}}
	pessoa2 := pessoa{
		nome:               "Natasha",
		sobrenome:          "Romanov",
		sorvetes_favoritos: [3]string{"Pistache", "Yogurte com morango", "Torta Holandesa"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.sorvetes_favoritos {
		fmt.Println("\t", v)
	}

	fmt.Println()

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.sorvetes_favoritos {
		fmt.Println("\t", v)
	}

}
