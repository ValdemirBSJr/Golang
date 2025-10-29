package main

import (
	"fmt"
)

func main() {

	sabores := []string{"peperoni", "mozarella", "abacaxi", "quatro_queijos", "marguerita"}

	fatia := sabores[0:2] // [:2]

	fmt.Println(fatia)

	fatia2 := sabores[2:5] //da 2 até o final [2:]

	fmt.Println(fatia2)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	//deletando pizza de abacaxi (sacrilegio)
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

}
