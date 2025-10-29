/*
1 - Ache o modulo no package.go.dev
2 - coloque o import no corpo do codigo
3 - Adicione os novos requisitos e hashes ao módulo:
4 - go mod init main.go
5 - go get sc.io/quote
5 - go mod tidy
*/

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
