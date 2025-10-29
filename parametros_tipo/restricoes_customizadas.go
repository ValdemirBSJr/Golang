/*
E se quisermos uma função Soma que funcione para uma fatia de ints e também para uma de float64s?
 A restrição comparable não nos ajuda aqui, pois ela não garante que o tipo suporte o operador +.

Para isso, podemos definir nossa própria restrição customizada usando uma interface.
*/

package main

import "fmt"

// Numero é uma restrição customizada que define um conjunto de tipos.
// A barra vertical '|' funciona como um "OU".
// Esta interface diz que qualquer tipo que satisfaça a restrição Numero
// deve ser, na verdade, um dos tipos listados: int, int64 ou float64.
type Numero interface {
	int | int64 | float64
}

// Soma calcula a soma de todos os valores em uma fatia.
// A declaração [T Numero] significa que o tipo 'T' DEVE ser
// um dos tipos definidos na nossa interface 'Numero'.
// O compilador agora sabe que qualquer 'T' aqui suportará o operador '+'.
func Soma[T Numero](valores []T) T {
	var total T // 'total' é inicializado com o valor zero do tipo T (0 para int, 0.0 para float)

	for _, valor := range valores {
		total += valor
	}
	return total
}

func main() {
	// Usando a função Soma com uma fatia de int64
	fatiaDeInteiros := []int64{10, 20, 30, 40}
	totalInteiros := Soma(fatiaDeInteiros)
	fmt.Printf("Soma dos inteiros: %d (Tipo: %T)\n", totalInteiros, totalInteiros)

	// Usando a MESMA função Soma com uma fatia de float64
	fatiaDeFloats := []float64{1.1, 2.2, 3.3, 4.4}
	totalFloats := Soma(fatiaDeFloats)
	fmt.Printf("Soma dos floats: %f (Tipo: %T)\n", totalFloats, totalFloats)

	// O código abaixo NÃO compila, pois 'string' não está na lista de tipos da interface 'Numero'.
	// fatiaDeStrings := []string{"a", "b", "c"}
	// Soma(fatiaDeStrings)

}
