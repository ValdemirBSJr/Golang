/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import "fmt"

func main() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "Doce",
		ondetem: []string{"Holanda", "Shopping", "Na casa do tio rico"},
		vaibemcom: map[string]string{
			"café da manhã": "Chá",
			"Almoço":        "Cafezinho",
		},
	}

	fmt.Println(x)

}
