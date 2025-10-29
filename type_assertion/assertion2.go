package main

import "fmt"

type Leitor interface {
	Ler() string
}

type Escritor interface {
	Escrever(string)
}

type LeitorEscritor interface {
	Leitor
	Escritor
}

type Arquivo struct {
	nome string
}

func (a Arquivo) Ler() string {
	return "lendo " + a.nome
}

func (a Arquivo) Escrever(dados string) {
	fmt.Println("Escrevendo", dados, "em", a.nome)
}

func main() {
	var leitor Leitor = Arquivo{nome: "O Alquimista.pdf"}

	// Verifica se o valor em 'leitor' também implementa a interface LeitorEscritor
	if le, ok := leitor.(LeitorEscritor); ok {
		fmt.Println("O valor implementa LeitorEscritor.")
		le.Escrever("anotações de dados importantes")
		fmt.Println(le.Ler())
	} else {
		fmt.Println("O valor não implementa LeitorEscritor.")
	}
}
