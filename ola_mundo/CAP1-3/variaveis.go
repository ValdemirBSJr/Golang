// o gopher := é pra atribuir dinamicamente o tipo da variável de acordo com o valor. Só funciona dentro de codeblocks
// x = 20 só pode por que a variável já existe e recebe um novo valor. Pra variaveis novas tem que ser :=
// posso usar em conjunto com variaveis existentes, desde que acrescente uma nova
// uma vez atribuido o tipo da variavel ela nao pode ser mudado. pra mudar tem que declarar novamente
package main

import (
	"fmt"
)

var z string = "boa tarde"
var k = "!"
var q string

var a int
var b float64
var c bool
var d string

func main() {
	x := 10
	y := "Olá, bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20

	q = "!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Println(z, q)

	x, w := 30, "foo"

	fmt.Println(x, w, k)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	string_literal := "Olá\nMundo!"
	string_crua := `olá\nmundo!`

	fmt.Println(string_literal)
	fmt.Println(string_crua)

	string_doida := `Olha 
	      a string
		  
		                    doida!`
	fmt.Println(string_doida)

}
