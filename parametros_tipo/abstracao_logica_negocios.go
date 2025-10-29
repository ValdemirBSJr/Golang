/*
Você pode usar generics para escrever a lógica central de um sistema (como um
pipeline de processamento de dados) de forma genérica, e depois fornecer as implementações
específicas para cada tipo de dado.

O que ganhamos aqui? Uma enorme redução na duplicação de código.
A lógica do "pipeline" foi escrita uma vez. Se amanhã você criar um tipo Produto que
também seja ItemProcessavel, ele funcionará imediatamente com ProcessarDados sem nenhuma alteração.
*/

package main

import "fmt"

// ItemProcessavel é uma interface que nossos dados devem satisfazer.
type ItemProcessavel interface {
	ID() string
}

// ProcessarDados itera sobre uma fatia de qualquer tipo que satisfaça a
// interface ItemProcessavel e aplica uma lógica comum.
func ProcessarDados[T ItemProcessavel](itens []T) {
	fmt.Println("-- Inciando pipeline de processamento ---")
	for _, item := range itens {
		fmt.Printf("Processando item com ID: %s\n", item.ID())
		// ... aqui poderia haver uma lógica complexa: salvar no banco,
		// enviar para uma fila, etc.
	}
	fmt.Println("--- Pipeline finalizado ---")
}

// Tipos de dados concretos que satisfazem a interface.
type Pedido struct{ Numero string }

func (p Pedido) ID() string { return p.Numero }

type NotaFiscal struct{ ChaveDeAcesso string }

func (nf NotaFiscal) ID() string { return nf.ChaveDeAcesso }

func main() {
	pedidos := []Pedido{{Numero: "P001"}, {Numero: "P002"}}
	notasFiscais := []NotaFiscal{{ChaveDeAcesso: "NF456"}, {ChaveDeAcesso: "NF789"}}

	// Usamos a mesma função para processar tipos de dados completamente diferentes.
	ProcessarDados(pedidos)
	ProcessarDados(notasFiscais)
}
