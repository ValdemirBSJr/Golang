package main

import (
	"fmt"
)

func main() {
	x := 2
	y := x >> 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)

	a := 24
	b := a << 2
	c := a >> 2

	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
}
