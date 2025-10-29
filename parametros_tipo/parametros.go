//parametros de tipo tambem são conhecidos como generics.
//Ele cria uma única função que funciona com fatias (slices)
//de diferentes tipos, desde que esses tipos possam ser comparados.

//os generics servem para muito mais do que comparações e cálculos.
//Eles são uma ferramenta fundamental para escrever código que é:

//Reutilizável: Escreva estruturas de dados e algoritmos uma vez.

//Abstrato: Separe a lógica do "como fazer" (o algoritmo genérico) dos "o quê fazer com"
//(os tipos concretos).

//Seguro: Mantenha toda a verificação de tipos em tempo de compilação que faz o Go ser tão robusto.

package main

import "fmt"

// Indice retorna a posição de 'elemento' em 'fatia', ou -1 se não for encontrado.
// [T comparable] é a declaração do parâmetro de tipo.
// - 'T' é um nome de placeholder para um tipo qualquer (poderia ser qualquer outra letra).
// - 'comparable' é uma restrição: diz que qualquer tipo usado no lugar de 'T'
//   DEVE suportar os operadores de comparação '==' e '!='.

func Indice[T comparable](fatia []T, elemento T) int {
	for i, valor := range fatia {
		// 'v' e 'elemento' são do tipo 'T'. Como 'T' tem a restrição 'comparable',
		// podemos usar o operador '==' com segurança aqui. O compilador garante isso.
		if valor == elemento {
			return i
		}
	}
	return -1
}

func main() {
	// Cria uma fatia de inteiros.
	fatiaDeInteiros := []int{10, 20, 15, -10}

	// Ao chamar Indice, Go infere automaticamente que T é 'int'.
	// 'int' é um tipo comparável, então a chamada é válida.
	fmt.Println("Posição do número 15:", Indice(fatiaDeInteiros, 15))

	// A mesma função Indice agora funciona com uma fatia de strings.
	fatiaDeStrings := []string{"go", "rust", "zig", "python"}

	// Desta vez, Go infere que T é 'string'.
	// 'string' também é um tipo comparável, então a chamada é válida.
	fmt.Println("Posição da string 'zig':", Indice(fatiaDeStrings, "zig"))
	fmt.Println("Posição da string 'java':", Indice(fatiaDeStrings, "java")) // retorna -1 nao tem

}
