/*
Este é, talvez, o caso de uso mais clássico. Antes,
se você quisesse uma estrutura de Pilha (Stack), precisaria de uma PilhaDeInts,
PilhaDeStrings, etc. Agora, você pode criar uma Pilha[T] que funciona para qualquer tipo.
*/

package main

import "fmt"

// Pilha é uma estrutura de dados genérica do tipo LIFO (Last-In, First-Out).
// O [T any] significa que esta Pilha pode armazenar elementos de qualquer tipo.

type Pilha[T any] struct {
	elementos []T
}

// Push adiciona um elemento do tipo T ao topo da pilha.
func (p *Pilha[T]) Push(elemento T) {
	p.elementos = append(p.elementos, elemento)
}

// Pop remove e retorna o elemento do topo da pilha.
// Retorna o elemento e um booleano indicando se a operação foi bem-sucedida.
func (p *Pilha[T]) Pop() (T, bool) {
	if len(p.elementos) == 0 {
		var zero T // Cria o "valor zero" para o tipo T (0, "", nil, etc.)
		return zero, false
	}
	ultimoIndice := len(p.elementos) - 1
	elemento := p.elementos[ultimoIndice]
	p.elementos = p.elementos[:ultimoIndice]
	p.elementos = p.elementos[:ultimoIndice]
	return elemento, true

}

// Definição de um tipo customizado para o exemplo.
type Usuario struct {
	ID   int
	Nome string
}

func main() {
	// --- Usando a Pilha com 'int' ---
	pilhaDeNumeros := Pilha[int]{}
	pilhaDeNumeros.Push(10)
	pilhaDeNumeros.Push(20)
	v, _ := pilhaDeNumeros.Pop()
	fmt.Printf("Elemento removido da pilha de números: %d\n", v) // Saída: 20

	// --- Usando a MESMA Pilha com 'Usuario' ---
	pilhaDeUsuarios := Pilha[Usuario]{}
	pilhaDeUsuarios.Push(Usuario{ID: 1, Nome: "Alice"})
	pilhaDeUsuarios.Push(Usuario{ID: 2, Nome: "Bob"})
	u, _ := pilhaDeUsuarios.Pop()
	fmt.Printf("Elemento removido da pilha de usuários: %+v\n", u) // Saída: {ID:2 Nome:Bob}
}
