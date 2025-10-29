/*
Mapas (Maps) em Go são o análogo direto dos Dicionários (Dictionaries) em Python.
Eles são coleções de pares chave-valor.

Porem são estaticas. Não adimitem elementos de tipos diferentes como no python
*/

package main

import "fmt"

func main() {
	notas := map[string]float64{"joao": 9.5, "maria": 8.0}

	// Tentando pegar a nota de Maria
	notaMaria, existeMaria := notas["maria"]
	if existeMaria {
		fmt.Printf("A nota de Maria é %.1f\n", notaMaria) // Saída: A nota de Maria é 8.0
	}

	// Tentando pegar a nota de Pedro, que não existe
	notaPedro, existePedro := notas["pedro"]
	if !existePedro {
		fmt.Println("Pedro não está no mapa.")                     // Saída: Pedro não está no mapa.
		fmt.Printf("O valor zero para float64 é: %v\n", notaPedro) // Saída: O valor zero para float64 é: 0
	}

	// Forma idiomática para verificar existência e obter valor ao mesmo tempo.
	// As variáveis 'notaMaria' e 'ok' só existem dentro deste bloco 'if'.
	if notaMaria, ok := notas["maria"]; ok {
		// Este bloco só executa se a chave "maria" existir (ok == true).
		fmt.Printf("A nota de Maria é %.1f\n", notaMaria)
	} else {
		// Opcional: bloco que executa se a chave não existir.
		fmt.Println("Maria não está no mapa.")
	}

	// Fazendo o mesmo para Pedro.
	if notaPedro, ok := notas["pedro"]; !ok {
		// Este bloco só executa se a chave "pedro" NÃO existir (ok == false).
		fmt.Println("Pedro não está no mapa.")
		fmt.Printf("O valor retornado (valor zero de float64) foi: %v\n", notaPedro)
	}

	// Cria um mapa vazio onde as chaves são strings e os valores são inteiros.
	// Se você não fizesse isso, o mapa seria `nil` e tentar adicionar uma chave causaria um `panic`.
	idades := make(map[string]int)
	idades["ana"] = 30
	idades["pedro"] = 25

	fmt.Println(idades)
}
