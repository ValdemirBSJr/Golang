/*
Os generics são perfeitos para criar funções utilitárias que operam sobre fatias e mapas,
como Map, Filter e Reduce.

A função Map (ou Mapear), por exemplo, transforma uma fatia de
um tipo T em uma fatia de outro tipo U, aplicando uma função.

O que ganhamos aqui? A capacidade de criar funções utilitárias poderosas e
abstratas. A lógica de "iterar e transformar" foi escrita uma vez e pode ser usada para
converter qualquer tipo de fatia em qualquer outra.

é como se passasse uma lista e retornasse un dicionario
*/
package main

import (
	"fmt"
	"strconv"
)

// Mapear transforma uma fatia de tipo T em uma fatia de tipo U.
//
// [T, U any]: Temos dois parâmetros de tipo, T (tipo de entrada) e U (tipo de saída).
// fn func(T) U: O segundo argumento é uma função que sabe como converter um T em um U.
func Mapear[T, U any](fatia []T, fn func(T) U) []U {
	resultado := make([]U, len(fatia))
	for i, v := range fatia {
		resultado[i] = fn(v)
	}
	return resultado
}

func main() {
	numeros := []int{1, 2, 3, 4}

	// Usamos Mapear para transformar uma fatia de int em uma fatia de string.
	// T é int, U é string.
	stringsFormatadas := Mapear(numeros, func(n int) string {
		return "Valor: " + strconv.Itoa(n)
	})

	fmt.Println("Fatia original (ints):", numeros)
	fmt.Println("Fatia transformada (strings):", stringsFormatadas)
}
