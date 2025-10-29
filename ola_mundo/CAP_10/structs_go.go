/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos diferentes.
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.
*/

package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := cliente{
		nome:      "Joao",
		sobrenome: "Da Silva",
		fumante:   false,
	}

	cliente2 := cliente{
		nome:      "Joana",
		sobrenome: "Pereira",
		fumante:   true,
	}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
	fmt.Println(cliente1.nome)

}
