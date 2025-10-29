/*
Go não tem classes. No entanto, você pode definir métodos em tipos.
O método é uma função com um argumento receptor especial.
O receptor aparece em sua lista de argumentos entre a própria
palavra-chave func e o nome do método.
*/

package main

import "fmt"

type Medidas struct {
	Largura, Altura float64
}

func (v Medidas) Area() float64 {
	return v.Largura * v.Altura
}

func main() {
	v := Medidas{3, 4}
	fmt.Printf("A área das medidas passadas é: %v m²\n", v.Area())
}
