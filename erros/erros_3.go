/*
Às vezes, um erro de uma função de baixo nível não tem contexto suficiente.
Por exemplo, um erro "arquivo não encontrado" não diz por que o programa estava
tentando abrir aquele arquivo.

A partir do Go 1.13, podemos "embrulhar" (wrap) erros para adicionar contexto
sem perder a informação do erro original. Isso é feito com a diretiva %w na função fmt.Errorf.

errors.Is: Verifica se um erro na cadeia de erros é igual a um erro específico.
errors.Unwrap: Retorna o erro original que foi "embrulhado".
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

// lerConfig simula a leitura de um arquivo de configuração.
func lerConfig(caminho string) error {
	// Tenta abrir o arquivo
	arquivo, err := os.Open(caminho)
	if err != nil {
		// Embrulha o erro original com mais contexto.
		// Agora sabemos que o erro do os.Open aconteceu durante a leitura da config.
		return fmt.Errorf("falha ao ler arquivo de configuração '%s': %w", caminho, err)
	}
	defer arquivo.Close()

	// ...lógica para processar o arquivo...
	fmt.Println("Arquivo de configuração lido com sucesso.")
	return nil
}

func main() {
	err := lerConfig("config.toml.inexistente")
	if err != nil {
		// Erro completo recebido
		fmt.Println("Erro completo recebido:", err)

		// Podemos "desembrulhar" o erro para inspecionar a causa original e mostrar apenas a parte que criamos
		errOriginal := errors.Unwrap(err)
		fmt.Println("Causa original:", errOriginal)

		// E podemos verificar se a causa é um erro específico
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Verificação: O arquivo realmente não existe. Crie um arquivo de configuração.")
		}
	}
}
