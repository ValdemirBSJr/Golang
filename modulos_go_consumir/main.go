/*
go mod init main

vi go.mod

A sintaxe é replace [nome do módulo que você quer substituir] => [caminho relativo para o módulo local].

go mod tidy

go get exemplo/boas_vindas
*/
package main

import (
	"exemplo/boas_vindas"
	"fmt"
	"log"
)

func main() {

	mensagem, err := boas_vindas.Boas_Vindas("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mensagem)

}
