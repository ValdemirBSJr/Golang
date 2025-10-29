/*
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
*/

package main

import (
	"fmt"
)

func main() {
	valores := [5]int{1, 2, 3, 4, 5}

	for i, v := range valores {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", valores)

}
