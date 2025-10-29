package main

import (
	"fmt"
)

func main() {

	s := "Olá, brincadeira!"
	sb := []byte(s)

	for _, v := range sb {

		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println()

	for _, v := range s {

		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}
}
