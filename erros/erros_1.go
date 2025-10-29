/*
Em go não temos erros como try catch. Os erros são tratatados como retorno de valores de uma função.
Erros são valores e devem ser tratadas de forma explicita.
*/

package main

import (
	"fmt"
	"log"
	"strconv"
)

func converterTxtNumero(texto string) (int, error) {
	//Atoi ja retorna int e mensagem então nao precisa declarar variaveis
	return strconv.Atoi(texto)

}

func main() {
	/// Cenário 1: Sucesso
	textoValido := "123"
	numero, err := converterTxtNumero(textoValido)
	if err != nil {
		// Se um erro ocorrer aqui, é algo inesperado.
		// Usar log.Fatal encerra o programa e imprime o erro.
		log.Fatalf("Falha inesperada ao converter '%s': %v", textoValido, err)
	}
	fmt.Printf("Conversão bem-sucedida! O número é %d.\n", numero)

	fmt.Println("--------------------")

	// Cenário 2: Falha
	textoInvalido := "não sou um número"
	// Usamos o "blank identifier" (_) para ignorar o primeiro valor de retorno (o número),
	// pois só nos interessa o erro neste caso.
	_, err = converterTxtNumero(textoInvalido)
	if err != nil {
		// Aqui, o erro é esperado. Apenas informamos o usuário.
		fmt.Printf("Ocorreu um erro esperado ao converter '%s': %v\n", textoInvalido, err)
	}

}
