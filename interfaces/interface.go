/*
Principais Conceitos para Lembrar:
Contrato de Comportamento: Uma interface em Go é um conjunto de assinaturas de métodos.
Ela define o que um tipo deve ser capaz de fazer, não o que ele é.
No exemplo, o comportamento é "ser capaz de enviar uma mensagem".

Satisfação Implícita (Duck Typing): Você não precisa declarar explicitamente
que seu tipo implementa uma interface. Se seu tipo tiver todos os métodos que
a interface exige (com as mesmas assinaturas), ele automaticamente a satisfaz.
“Se anda como um pato e grasna como um pato, então é um pato.”

Polimorfismo: Interfaces permitem que você escreva funções que operam sobre diferentes tipos
de dados de maneira uniforme. A função DispararNotificacao é polimórfica:
ela funciona com Email, SMS e qualquer outro Notificador que você criar, sem precisar de
if/else ou switch para verificar o tipo. Isso torna seu código muito mais flexível e extensível.
*/
package main

import "fmt"

// --- 1. Definição da Interface ---
//
// Notificador é a nossa interface. Ela define um "contrato" de comportamento.
// Qualquer tipo que implementar o método `Enviar(mensagem string) error`
// será, implicitamente e automaticamente, considerado um tipo Notificador.
// Em Go, não usamos a palavra-chave "implements". A satisfação é implícita.
type Notificador interface {
	// Esta INTERFACE descreve o comportamento de "ser capaz de enviar".
	// Ela não diz COMO enviar, apenas QUE o tipo deve saber enviar.
	// É uma lista de requisitos.
	Enviar_mensagem(mensagem string) error
}

// --- 2. Criação de Tipos Concretos ---
//
// Agora, vamos criar tipos diferentes que podem se comportar como um Notificador.

// Email é um tipo concreto que usaremos para enviar notificações por email.
type Email struct {
	destinatario, assunto string
}

// SMS é outro tipo concreto, usado para enviar notificações por SMS.
type SMS struct {
	numeroTelefone string
}

// --- 3. Implementação da Interface pelos Tipos Concretos ---
//
// Para que os tipos Email e SMS satisfaçam a interface Notificador,
// eles precisam ter um método com a exata mesma assinatura de Enviar().

// Implementação do método Enviar para o tipo Email.
// O receptor `(e Email)` associa este método à struct Email.
// Como Email agora tem este método, ele satisfaz a interface Notificador.
// Ele define COMO um Email envia uma notificação.
// Uma interface, por outro lado, não pertence a nenhum tipo concreto e
// não tem código de implementação. Ela é uma definição abstrata de comportamento.
// É um contrato que diz: "Para ser considerado deste grupo (ex: Notificador),
// você precisa ser capaz de fazer estas coisas (ex: ter um método Enviar)".
func (e Email) Enviar_mensagem(mensagem string) error {
	// Simula o envio de um email.
	fmt.Printf("Enviando email para '%s' com assunto '%s': %s\n", e.destinatario, e.assunto, mensagem)
	return nil // `nil` em Go significa "sem erro".
}

// Implementação do método Enviar para o tipo SMS.
// O receptor `(s SMS)` associa este método à struct SMS.
// Agora, SMS também satisfaz a interface Notificador.
func (s SMS) Enviar_mensagem(mensagem string) error {
	fmt.Printf("Enviando SMS para %s: %s\n", s.numeroTelefone, mensagem)
	return nil
}

// --- 4. Usando a Interface (O Poder do Polimorfismo) ---
//
// Esta função aceita QUALQUER tipo que seja um Notificador.
// Ela não se importa se o valor concreto é um Email, um SMS, ou qualquer outro tipo
// que venha a ser criado no futuro, desde que ele saiba como se "Enviar".
// Este é o poder do polimorfismo em Go.

func DispararNotificacao(n Notificador, mensagem string) {
	fmt.Println("Disparando notificação...")
	err := n.Enviar_mensagem(mensagem) // O método correto (de Email ou SMS) é chamado aqui.

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}

// --- 5. Função Principal: Juntando Tudo ---
func main() {
	// Criando instâncias dos nossos tipos concretos.
	emailDoUsuario := Email{destinatario: "seuNinguem@exemplo.com.br", assunto: "cobrança por inadimplência"}
	smsUsuario := SMS{numeroTelefone: "83 3281-1098"}

	// Podemos passar tanto um Email quanto um SMS para a mesma função,
	// porque ambos satisfazem a interface Notificador.
	DispararNotificacao(emailDoUsuario, "Olá cara! me pague o aluguel!")
	DispararNotificacao(smsUsuario, "Maria, quer vender um rim?")

	fmt.Println("\n--- Trabalhando com uma lista de Notificadores ---")
	// Outro uso poderoso é criar coleções de interfaces.
	// A slice `notificadores` pode conter valores de tipos diferentes (Email, SMS),
	// desde que todos eles satisfaçam a interface Notificador.

	notificadores := []Notificador{
		Email{destinatario: "chefe@empresa.com", assunto: "Relatório Semanal"},
		SMS{numeroTelefone: "+55 (11) 77777-6666"},
		Email{destinatario: "financeiro@empresa.com", assunto: "Nota Fiscal"},
	}

	// Podemos iterar sobre a slice e tratar cada elemento da mesma forma,
	// independentemente do seu tipo concreto.
	for _, notificador := range notificadores {
		notificador.Enviar_mensagem("Este é um comunicado importante para todos.")
	}

} //main
