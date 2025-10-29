/*
O próximo passo depois de entender as goroutines é aprender como elas se comunicam.
Os canais (channels) são a resposta idiomática do Go para essa necessidade.
Eles são a principal ferramenta para garantir uma comunicação segura entre goroutines,
evitando condições de corrida (race conditions) e outros problemas comuns em programação concorrente.

"Não comunique compartilhando memória; em vez disso, compartilhe memória comunicando."

Pense em um canal como uma esteira rolante ou um duto que conecta duas ou mais goroutines.
Tipado: Você declara o tipo de dado que pode passar por aquele canal
(ex: chan int para inteiros, chan string para strings).

Seguro para Concorrência: O Go garante que apenas uma goroutine possa acessar o
dado no canal por vez. Você não precisa se preocupar com locks ou mutexes para operações simples.

Sincronização: Enviar um dado para um canal ou receber um dado de um canal
são operações que bloqueiam a execução da goroutine até que a outra ponta
(o receptor ou o emissor) esteja pronta. Essa é a sua característica mais poderosa.
*/
package main

import "fmt"

// `somar` é uma função que calcula a soma de uma fatia de inteiros.
// Ela recebe a fatia de números e um canal para onde enviará o resultado.
// A sintaxe `chan int` significa "um canal de inteiros".
func somar(fatia []int, canal chan int) {
	// Cria uma variável local para guardar a soma.
	soma := 0
	// Itera sobre cada valor `v` dentro da `fatia`.
	for _, v := range fatia {
		soma += v
	}

	// O OPERADOR DE CANAL `<-`
	// A sintaxe `canal <- soma` significa "envie o valor da variável `soma`
	// PARA DENTRO do `canal`".
	//
	// **ESTA OPERAÇÃO É BLOQUEANTE!**
	// A goroutine que está executando esta função irá pausar nesta linha
	// até que alguma outra goroutine esteja pronta para RECEBER o valor do canal.
	canal <- soma
}

func main() {
	// Nossa fatia de números que queremos somar.
	numeros := []int{7, 2, 8, -9, 4, 0}

	// A FUNÇÃO `make`:
	// Criamos um canal usando `make(chan int)`.
	// Agora temos um "duto" pronto para transportar inteiros entre goroutines.
	canal := make(chan int)

	// Dividimos o trabalho:
	// 1. Lançamos uma primeira goroutine para somar a primeira metade da fatia.
	//    A sintaxe `numeros[:len(numeros)/2]` pega os elementos do início até a metade.
	//    Passamos o mesmo canal para ela.
	go somar(numeros[:len(numeros)/2], canal)

	// 2. Lançamos uma segunda goroutine para somar a segunda metade da fatia.
	//    A sintaxe `numeros[len(numeros)/2:]` pega os elementos da metade até o fim.
	//    Passamos o MESMO canal para ela também.
	go somar(numeros[len(numeros)/2:], canal)

	// RECEBENDO VALORES DO CANAL:
	// A sintaxe `<-canal` significa "receba um valor DE DENTRO do `canal`".
	//
	// **ESTA OPERAÇÃO TAMBÉM É BLOQUEANTE!**
	// A goroutine `main` irá pausar aqui na primeira recepção (`<-canal`)
	// até que uma das goroutines `somar` envie um valor.
	// Depois, ela continuará e pausará novamente na segunda recepção, esperando
	// o valor da outra goroutine `somar`.
	primeiraSoma, segundaSoma := <-canal, <-canal

	// Após receber os dois valores, a goroutine `main` é desbloqueada e pode continuar.
	// Imprimimos as somas parciais e a soma total.
	fmt.Println("Primeira soma parcial:", primeiraSoma)
	fmt.Println("Segunda soma parcial:", segundaSoma)
	fmt.Println("Soma total:", primeiraSoma+segundaSoma)
}
