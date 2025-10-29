/*
O type switch é um complemento natural ao que estávamos discutindo sobre interfaces.
Enquanto as interfaces nos permitem tratar diferentes tipos de forma uniforme (polimorfismo),
o type switch nos dá uma maneira segura e elegante de fazer o oposto: descobrir qual é o
tipo concreto específico que está armazenado dentro de uma variável de interface e agir de acordo.

O que é o type switch?
É uma construção da linguagem Go, similar a um switch normal, mas em vez de verificar o
valor de uma variável, ele verifica o tipo dela. Ele é usado quase que exclusivamente com variáveis
do tipo interface.

Seu principal objetivo é permitir que você execute um bloco de código diferente
para cada tipo concreto que uma interface possa conter.
*/

package main

import "fmt"

type Notificador interface {
	Enviar(mensagem string) error
}

type Email struct {
	assunto, destinatario string
}

func (e Email) Enviar(mensagem string) error {
	fmt.Printf("Enviando email para %s: %s\n", e.destinatario, mensagem)
	return nil
}

type SMS struct {
	numeroTelefone string
}

func (s SMS) Enviar(mensagem string) error {
	fmt.Printf("Enviando SMS para %s: %s\n", s.numeroTelefone, mensagem)
	return nil
}

// Adicionando um novo tipo para demonstrar a flexibilidade
type Push struct {
	idDispositivo string
}

func (p Push) Enviar(mensagem string) error {
	fmt.Printf("Enviando notificação Push para %s: %s\n", p.idDispositivo, mensagem)
	return nil
}

// --- A função que usa o type switch ---

// AnalisarNotificador recebe uma interface e imprime detalhes específicos
// de acordo com o tipo concreto que ela contém.
func AnalisarNotificador(n Notificador) {
	fmt.Println("--- Analisando Notificador ---")

	// Este é o type switch.
	// Ele verifica o tipo concreto da variável 'n'.
	switch v := n.(type) {
	case Email:
		// Dentro deste bloco, 'v' é uma variável do tipo Email.
		// Portanto, podemos acessar seus campos específicos como 'v.assunto'.
		fmt.Printf("Tipo: Email. Destinatário: %s, Assunto: %s\n", v.destinatario, v.assunto)
	case SMS:
		// Aqui, 'v' é do tipo SMS.
		// Podemos acessar o campo 'v.numeroTelefone'.
		fmt.Printf("Tipo: SMS. Número de Telefone: %s\n", v.numeroTelefone)

	case Push:
		// E aqui, 'v' é do tipo Push.
		fmt.Printf("Tipo: Push. ID do Dispositivo: %s\n", v.idDispositivo)

	default:
		// Este bloco é executado se 'n' contiver um tipo
		// que não listamos nos casos acima.
		fmt.Println("Tipo: Notificador desconhecido.")
	}

	fmt.Println("------------------------------")
}

func main() {
	// Criamos uma slice de Notificadores, contendo tipos concretos diferentes.
	notificadores := []Notificador{
		Email{destinatario: "ana@example.com", assunto: "Seu pedido foi enviado"},
		SMS{numeroTelefone: "+55 11 98765-4321"},
		Push{idDispositivo: "APA91bHun4..."},
	}

	// Iteramos sobre a slice e usamos nossa função com o type switch.
	for _, n := range notificadores {
		// A função AnalisarNotificador executará um código diferente para cada tipo.
		AnalisarNotificador(n)
	}
}
