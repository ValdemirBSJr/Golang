package main

import (
	"fmt"
)

var x bool

func main() {

	fmt.Println(x) //valor padrao
	x = true
	fmt.Println(x) //atribuido
	x = (10 > 100)
	fmt.Println(x)

	y := (10 < 100)
	fmt.Println(y)

}
