// crie um loop que mostre os anos desde q vc nasceu
package main

import (
	"fmt"
)

func main() {

	anoemquenasci := 1984
	anocorrente := 2024

	for anoemquenasci <= anocorrente {
		fmt.Println(anoemquenasci)
		anoemquenasci++
	}

}
