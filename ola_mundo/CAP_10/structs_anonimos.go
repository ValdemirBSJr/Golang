// structs anonimos usa uma vez e descarta
package main

import (
	"fmt"
)

func main() {
	//a identacao tem que ser exatamente como abaixo ou da erro. Sem virgula final e com a chave fechada junto ao 50
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Maria",
		idade: 50}

	fmt.Println(x)
}
