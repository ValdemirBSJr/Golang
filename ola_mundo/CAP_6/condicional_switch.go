package main

import (
	"fmt"
)

var x interface{}

func main() {

	x = true

	switch x.(type) {
	case int:
		fmt.Println("É um int")
	case bool:
		fmt.Println("É um bool")
	case string:
		fmt.Println("É um texto")
	case float64:
		fmt.Println("É um float")
	}

	switch y := 2; {
	case y == 1:
		fmt.Println("É 1")
	case y == 2:
		fmt.Println("É 2")
	case y == 3:
		fmt.Println("É 3")
	case y == 4:
		fmt.Println("É 4")
	}

}
