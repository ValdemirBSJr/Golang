package main

import (
	"fmt"
	"math/rand" //gerar numeros aleatorios
	"time"      //lib de manipular tempo
)

func gerar_senha_aleatoria(comprimento int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	password := make([]byte, comprimento)
	for i := range password {
		password[i] = charset[random.Intn(len(charset))]
	}
	return string(password)
}

func main() {
	password := gerar_senha_aleatoria(12)
	fmt.Printf("Senha gerada: %s\n", password)
}
