package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[2][1])
}
