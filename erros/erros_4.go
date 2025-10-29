/*
Para erros mais complexos, você pode definir seu próprio tipo de erro.
Isso permite anexar informações extras (como um código de status HTTP, por exemplo).
Para verificar e extrair esse tipo de erro customizado da cadeia, usamos errors.As.

O asterisco * no método define um "pointer receiver", significando que o método opera no
valor original na memória, não em uma cópia.
Analogia: Você entrega o documento original para alguém.
Qualquer alteração feita nele é permanente.
Vantagem: É mais eficiente para structs grandes e é a única maneira de um método
poder modificar o estado da struct na qual ele opera.
No nosso caso: Usamos um pointer receiver porque é a convenção
para tipos que implementam interfaces (como a interface error).
Embora não estejamos modificando a struct MeuErroDeRede dentro do método Error(),
usar ponteiros é mais performático e consistente.

O operador & age sobre o resultado do passo 1. Ele diz: "Não me dê o valor
da struct que você acabou de criar. Em vez disso, me dê o endereço de memória
onde essa nova struct foi armazenada".

RESUMINDO
return &MeuErroDeRede{...}: Cria uma nova struct MeuErroDeRede na memória
e retorna o endereço dela. O valor retornado é do tipo *MeuErroDeRede (um ponteiro).

func (e *MeuErroDeRede) Error(): Define um método chamado Error que pode ser chamado por
qualquer variável que seja um ponteiro para uma MeuErroDeRede (como a variável err que recebeu
o retorno da outra função).
*/
package main

import (
	"errors"
	"fmt"
)

// MeuErroDeRede é um tipo de erro customizado.
// Pense nisso como um "molde" para guardar informações
// detalhadas sobre um erro específico que pode acontecer na sua aplicação.
// Por si só, isso é apenas uma struct comum, um contêiner de dados.
type MeuErroDeRede struct {
	URL    string
	Codigo int
	Msg    string
}

// Implementamos ao metodo 'error'
// A linguagem Go tem uma interface pré-definida chamada 'error'.
// Ela é extremamente simples:
//
//	type error interface {
//	    Error() string
//	}
//
// Qualquer tipo que implemente um método com a assinatura exata "Error() string"
// automaticamente satisfaz essa interface. Ao fazer isso, você está dizendo ao Go:
// "Ei, essa minha struct 'MeuErroDeRede' agora também PODE SER TRATADA COMO um erro padrão".
func (e *MeuErroDeRede) Error() string {
	// Este método define QUAL TEXTO APARECERÁ quando alguém tentar
	// "imprimir" este erro (como o fmt.Println faz).
	// Portanto, este método SERÁ SIM chamado. Veremos como na função main.
	return fmt.Sprintf("Falha na operação de rede para %s (%d): %s", e.URL, e.Codigo, e.Msg)
}

// simularFalha simula uma chamada de API que falha.
// Note que o tipo de retorno da função é 'error', a interface.
func simularFalha() error {
	// Retorna uma instância do nosso erro customizado.
	// Como *MeuErroDeRede implementa o método Error(), ele satisfaz a interface 'error'.
	// Isso significa que podemos retornar uma instância de &MeuErroDeRede
	// em uma função que promete retornar um 'error'.
	// O Go permite isso, pois nosso tipo agora "é um" erro.
	return &MeuErroDeRede{
		URL:    "https://api.example.com/data",
		Codigo: 503,
		Msg:    "Serviço indisponível",
	}
}

func main() {
	// A variável 'err' aqui é do tipo 'error' (a interface),
	// mas o valor concreto que ela armazena é um ponteiro para a nossa
	// struct: *MeuErroDeRede.
	err := simularFalha()
	if err != nil {
		// QUANDO esta linha é executada, o pacote 'fmt' é inteligente.
		// Ele verifica se a variável 'err' implementa a interface 'error'.
		// Como a resposta é SIM, em vez de imprimir os campos da struct,
		// o 'fmt.Println' AUTOMATICAMENTE CHAMA O MÉTODO .Error()
		// para obter a representação em string do erro.
		// É por isso que você não chama err.Error() explicitamente, mas ele é chamado!
		fmt.Println("Ocorreu um erro:", err)

		// Agora, queremos acessar os campos específicos da nossa struct (*MeuErroDeRede),
		// como 'Codigo' e 'URL'. Mas a variável 'err' é do tipo 'error', que não tem esses campos.
		// Precisamos verificar se o erro que recebemos é do nosso tipo customizado.

		// Declara uma variável do tipo do nosso erro, inicialmente vazia.
		var erroDeRede *MeuErroDeRede
		// errors.As verifica se 'err' (ou qualquer erro embrulhado dentro dele)
		// é do tipo *MeuErroDeRede. Se for, atribui o valor a erroDeRede e retorna true.
		// errors.As é a forma moderna e segura de fazer uma "conversão de tipo" para erros.
		// Ele tenta "despejar" o valor de 'err' na variável 'erroDeRede'.
		if errors.As(err, &erroDeRede) {
			// Se chegamos aqui, significa que a conversão deu certo!
			// 'err' realmente era um *MeuErroDeRede.
			// Agora a variável 'erroDeRede' contém o valor e podemos acessar seus campos.
			fmt.Println("--- Detalhes do Erro de Rede ---")
			fmt.Printf("Código de Status: %d\n", erroDeRede.Codigo)
			fmt.Printf("URL: %s\n", erroDeRede.URL)
			fmt.Println("Ação sugerida: Tente novamente mais tarde ou verifique o status do serviço.")
		}
	}
}
