package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 5555124,
		"Joana":   999674,
	}

	amigos["Gopher"] = 4444

	fmt.Println(amigos)
	fmt.Println(amigos["Joana"])
	fmt.Println(amigos["Gopher"])
	// para saber se existe ou nao, usa o comma ok. Ele retorna o valor da chave e um booleano
	// o booleano so retorna true se existir

	sera, comma_ok := amigos["fantasma"]

	fmt.Println(sera, comma_ok)

	if sera, comma_ok := amigos["fantasma"]; !comma_ok {
		fmt.Println("Esse cara não é seu amigo")
	} else {
		fmt.Println("Esse cara é seu amigo ", sera)
	}

	amigos["Anitta"] = 696969
	fmt.Println(amigos)
	delete(amigos, "Anitta")
	fmt.Println(amigos)

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	}

	for k, v := range qualquercoisa {
		fmt.Println(k, v)
	}

}
