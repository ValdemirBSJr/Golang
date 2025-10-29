/*
mkdir modulos_go
cd modulos_go

go mod init exemplo/boas_vindas
vi boas_vindas.go

*/

package boas_vindas

import (
	"errors"
	"fmt"
)

func Boas_Vindas(nome string) (string, error) {
	if nome == "" {
		return "", errors.New("O nome não pode estar em branco")
	}
	mensagem := fmt.Sprintf("Olá, %v! Bem vindo!", nome)
	return mensagem, nil
}
