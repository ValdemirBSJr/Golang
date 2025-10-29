/*
Fatias (Slices) em Go são o análogo mais próximo das Listas (Lists) em Python.
Elas são coleções ordenadas e indexadas de elementos.

Porem são estaticas. Não adimitem elementos de tipos diferentes como no python

tem len(quantos itens a fatia atualmente contém.) e cap(capacidade total,
sem precisar de nova alocação e memória)
*/

package main

import "fmt"

func main() {
	// Cria uma fatia com 0 elementos, mas com espaço pré-alocado para 10.
	// len=0, cap=10.
	fatia := make([]int, 0, 10)
	// Agora podemos usar `append` até 10 vezes sem que Go precise alocar um novo array.
	fatia = append(fatia, 1, 2, 3) // len=3, cap=10

	fmt.Println(fatia)

}
