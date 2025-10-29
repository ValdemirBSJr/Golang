// ira somar/arrumar os modulos: go mod tidy (caso vc nao tenha instalado o modulo, ele vai ler o import do arquivo e instala)
// instalar um pacote diretamente: go get rsc.io/quote
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	numerodebytes, erros := fmt.Println("Olá", "gente!", 100)
	// se nao preciso de uma dessas, coloca o _ -> _, erros := fmt.Println("Olá", "gente!", 100)
	fmt.Println(quote.Go())
	fmt.Println(numerodebytes, erros)
}
