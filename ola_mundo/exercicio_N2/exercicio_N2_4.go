// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
package main

import (
	"fmt"
)

func main() {
	x := 200
	fmt.Printf("%d, %b, %#x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d, %b, %#x", y, y, y)

}
