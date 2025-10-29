/*
https://github.com/Ali-aqrabawi/gomiko

go mod init teste_gomiko
go get -u github.com/Ali-aqrabawi/gomiko/pkg
*/
package main

import (
	"fmt"
	"log"

	gomiko "github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	logarTimeout()
}

func logarTimeout() {
	timeout := gomiko.TimeoutOption(10)
	dispositivo, err := gomiko.NewDevice("187.64.56.62", "N5669203", "Val,.2510", "cisco_ios", 22, timeout)
	if err != nil {
		log.Fatal(err)
	}

	//Abrindo sessao
	if err := dispositivo.Connect(); err != nil {
		log.Fatal(err)
	}

	//envia o comando
	saida1, _ := dispositivo.SendCommand("show clock")

	//envia conf
	//comandos := []string{"comando 1"", "comando 2"}
	//saida2, _ := dispositivo.SendConfigSet(comandos)

	dispositivo.Disconnect()

	fmt.Println(saida1)
	//fmt.Println(saida2)
}
