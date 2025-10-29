/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/
package main

import (
	"fmt"
)

type pessoa struct {
	nome               string
	sobrenome          string
	sorvetes_favoritos []string //pode ser assim tbm []string e fica em aberto o tamanho
}

func main() {

	meumapa := make(map[string]pessoa)

	meumapa["Pimentao"] = pessoa{
		nome:               "Renata",
		sobrenome:          "Pimentão",
		sorvetes_favoritos: []string{"pistache", "morango", "baunilha"},
	}

	meumapa["da Prússia"] = pessoa{
		"Frederico",
		"da Prussia",
		[]string{"Chocolate", "Açaí"},
	}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é:", valor.nome, "e meus sorvetes favoritos são:")
		for _, sabor := range valor.sorvetes_favoritos {
			fmt.Println("-", sabor)
		}
		fmt.Println()
	}

}
