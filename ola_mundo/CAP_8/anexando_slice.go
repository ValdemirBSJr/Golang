package main

import (
	"fmt"
)

func main() {

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)

	//os 3 pontinhos indicam q vc passou os elementos de uma slice e nao a slice em si. Se nao tiver eles da erro
	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)

}
