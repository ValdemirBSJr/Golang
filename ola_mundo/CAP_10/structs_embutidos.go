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

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {

	pessoa1 := pessoa{nome: "Alfredo", idade: 30}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Janja",
			idade: 40,
		},
		titulo:  "Primeira dama do Brasil",
		salario: 80000,
	}

	pessoa3 := pessoa{"Mauricio", 45}

	pessoa4 := profissional{pessoa{"Carlos", 53}, "Aposentado folgado", 1500}

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.pessoa)
	fmt.Println(pessoa2.pessoa.idade)
	//se nao tiver colisao de nomes, posso abreviar o acima
	fmt.Println(pessoa2.idade)
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)

}
