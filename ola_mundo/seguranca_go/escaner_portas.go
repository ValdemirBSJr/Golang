package main

import (
	"fmt"
	"net"     //funcionalidades de rede
	"strconv" //converter cadeias de caracteres e tipos de dados numerios
	"time"    //funcionalidades de tempo
)

func main() {
	alvo := "testhtml5.vulnweb.com"

	for porta := 1; porta <= 1024; porta++ {
		endereco := alvo + ":" + strconv.Itoa(porta)               // concatena o nome do alvo(servidor) com a porta
		conn, err := net.DialTimeout("tcp", endereco, time.Second) //tenta  conexao TCP e usa o tempo limite de 1 seg

		//se der erro, se a porta nao responder vai pro proximo loop sem exibir mensagem
		if err != nil {
			continue
		}

		//se a conexao for estabelecida com sucesso, sera agendada para fechar no final da funcao
		defer conn.Close()
		//se for estabelecida, exibe mensagem
		fmt.Printf("Porta %d aberta\n", porta)

	}
}
