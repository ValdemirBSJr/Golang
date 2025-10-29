// https://pkg.go.dev/fmt
package main

import (
	"fmt"
)

func main() {

	x := "oi"
	y := "bom dia!"

	//Sprint serve pra juntar string e retornar para variaveis. Não imprime nada na tela
	z := fmt.Sprintln(x, y)

	fmt.Print(z)

}
