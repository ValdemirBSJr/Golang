/*IOTA SERVE PARA QUANDO VC PRECISAR DE VALORES CONSTANTES DIFERENTES
E NAO QUISER ATRIBUIR UM VALOR
_ pra jogar o valor fora e começar do 1
vc pode concatenar o iota para ele começar de um determinado valor (c)*/

package main

import (
	"fmt"
)

const (
	_ = iota
	a = iota
	b = iota
	x = iota
	y = iota
	z = iota
)

const (
	c = iota + 1000000
	d
	e
)

func main() {
	fmt.Println(a, b, x, y, z)
	fmt.Println(c, d, e)
}
