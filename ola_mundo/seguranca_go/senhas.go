package main

import (
	"fmt"
	"regexp" //trabalar com expressoes regulares
)

func isStrongPassword(senha string) bool {
	//verificar o tamanho minimo da string
	if len(senha) < 8 {
		return false
	}

	//verificar ao menos a presença de uma letra maiuscula
	if correspondencia, _ := regexp.MatchString(`[A-Z]`, senha); !correspondencia {
		return false
	}

	//verificar a presença de ao menos um numero
	if correspondencia, _ := regexp.MatchString(`[0-9]`, senha); !correspondencia {
		return false
	}

	//verifica a presença de caracteres especiais
	if correspondencia, _ := regexp.MatchString(`[$@!]`, senha); !correspondencia {
		return false
	}

	return true
}

func main() {
	senha := "p@ssw0rD" //senha que deseja verificar
	e_forte := isStrongPassword(senha)

	if e_forte {
		fmt.Println("A senha é forte e segura")
	} else {
		fmt.Println("A senha não cumpre o mínimo necessário de segurança")
	}
}
