/*
GO é fortemente tipada entao, podemos facilmente construir tipos para armazenar dados
isto é muito usado na linguagem
*/
package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 10

func main() {

	fmt.Printf("%v %T\n", b, b)

	//para transformar um tipo em outro (conversão)
	x := int(b)
	fmt.Printf("%v %T", x, x)

}
