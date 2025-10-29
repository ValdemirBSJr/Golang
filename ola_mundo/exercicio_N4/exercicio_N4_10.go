/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
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

	maps_cadastro["loureiro_kiko"] = []string{"Usar os trequinho no punho", "tacar fogo na guitarra"}

	delete(maps_cadastro, "pereira_luiz")

	for c, v := range maps_cadastro {
		fmt.Println(c)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}

	}

}
