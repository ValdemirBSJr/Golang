package main

import "fmt"

type Vortex struct {
	X, Y int
}

//forma mais explicita de declarar variaveis
var (
	v1       = Vortex{1, 2}
	v2       = Vortex{X: 1}
	ponteiro = &Vortex{1, 2}
)

func main() {
	fmt.Println(v1, ponteiro, v2)
}
