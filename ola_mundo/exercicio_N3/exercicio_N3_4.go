//Utilize um for "loop infinito" para demonstrar os anos desde que você nasceu.

package main

import (
	"fmt"
)

func main() {
	anoemquenasci := 1984
	anocorrente := 2024

	for {

		if anoemquenasci > anocorrente {
			break
		}

		fmt.Println(anoemquenasci)
		anoemquenasci++
	}
}
