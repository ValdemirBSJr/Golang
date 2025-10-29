/*
rodar o arquivo go: go run ola_mundo.go
compilar o arquivo .go: go build ola_mundo.go
*/
package main

import (
	"bufio" // Pacote necessário para ler a entrada do usuário
	"fmt"
	"os" // Pacote necessário para interagir com o sistema operacional
)

func main() {
	fmt.Println("Olá mundo!")

	fmt.Println("Pressione Enter para sair...")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
