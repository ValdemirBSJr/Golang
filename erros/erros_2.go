/*
Você deve projetar suas próprias funções para seguir o mesmo padrão.
Se sua função pode falhar, faça-a retornar um error.

Exemplo: Uma função de divisão que evita divisão por zero

Usamos os pacotes errors para criar um erro simples ou fmt para criar um erro formatado.
*/

package main

import (
	"errors"
	"fmt"
)

// dividir retorna o resultado da divisão ou um erro se o divisor for zero.
func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		// Cria um novo erro. errors.New é para erros estáticos simples.
		return 0, errors.New("não é possível dividir por zero")
	}

	return dividendo / divisor, nil // Retorna o resultado e um erro nulo (sucesso)
}

func main() {
	resultado, err := dividir(10.0, 2.0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado 1:", resultado)
	}

	resultado, err = dividir(10.0, 0.0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado 2:", resultado)
	}
}
