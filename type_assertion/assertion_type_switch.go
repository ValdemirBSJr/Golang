package main

import "fmt"

type Leitor interface {
	Ler() string
}

type Escritor interface {
	Escrever(string)
}

// Interface que combina as duas anteriores
type LeitorEscritor interface {
	Leitor
	Escritor
}

// --- Tipo concreto que implementa as interfaces ---
type Arquivo struct {
	nome string
}

func (a Arquivo) Ler() string {
	return "Lendo " + a.nome
}

func (a Arquivo) Escrever(dados string) {
	fmt.Println("Escrevendo", dados, "em", a.nome)
}

// --- Outro tipo para demonstração ---
type Console struct{}

func (c Console) Ler() string {
	return "Lendo da entrada padrão. Console"
}

// --- Função que utiliza o type switch ---
func descrever(l Leitor) {
	fmt.Println("--- Verificando Tipo ---")

	switch v := l.(type) {
	case LeitorEscritor:
		// Se o tipo concreto implementa LeitorEscritor
		fmt.Println("O valor é um LeitorEscritor.")
		fmt.Printf("Tipo concreto: %T\n", v)
		v.Escrever("anotações importantes de rodapé") // Podemos chamar métodos de Escritor
		fmt.Println(v.Ler())                          // e também métodos de Leitor

	case Console:
		// Se o tipo concreto é Console
		fmt.Println("O valor é apenas um Console Leitor.")
		fmt.Printf("Tipo concreto: %T\n", v)
		fmt.Println(v.Ler())

	case Arquivo:
		// Este caso não será atingido no nosso exemplo se o caso LeitorEscritor vier primeiro,
		// pois Arquivo satisfaz LeitorEscritor. A ordem dos 'case' importa!
		fmt.Println("O valor é um Arquivo que é apenas Leitor.")
		fmt.Printf("Tipo concreto: %T\n", v)
		fmt.Println(v.Ler())

	default:
		// Caso nenhum dos tipos acima corresponda
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
	fmt.Println("----------------------------------")
}

func main() {
	// arq implementa Leitor e Escritor, então satisfaz LeitorEscritor
	var arq Leitor = Arquivo{nome: "meu_arquivo.txt"}
	descrever(arq)

	// console só implementa Leitor
	var console Leitor = Console{}
	descrever(console)
}
