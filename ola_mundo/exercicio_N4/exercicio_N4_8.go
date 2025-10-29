/*
- Crie um map com key tipo string e value tipo []string.
  - Key deve conter nomes no formato sobrenome_nome
  - Value deve conter os hobbies favoritos da pessoa
*/
package main

import (
	"fmt"
)

func main() {

	maps_cadastro := map[string][]string{
		"albuquerque_carlos": []string{
			"Andar de bicicleta",
			"Caça esportiva",
		},
		"pereira_luiz": []string{
			"taxidermia",
			"Jetski",
		},
		"dosAnjos_jorge": []string{
			"Escarificação",
			"BAter em mendingo",
		},
	}

	for c, v := range maps_cadastro {

		fmt.Println(c, v)

	}

	for c, v := range maps_cadastro {
		fmt.Println(c)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}

	}

}
