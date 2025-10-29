/*
Em Go, a type assertion (asserção de tipo) é um mecanismo
que permite acessar o valor concreto de uma variável de interface. Essencialmente,
é uma forma de dizer ao compilador: "eu acredito que o valor armazenado nesta interface
é de um tipo específico e quero acessá-lo como tal".

valor := v.(T)
Neste caso, se a variável v realmente contiver um valor do tipo T, a operação
será bem-sucedida e o valor concreto será atribuído a valor.

Cuidado: Se v não contiver um valor do tipo T, o programa entrará em pânico
(um erro em tempo de execução que interrompe a execução normal do programa).

Para evitar o pânico, Go oferece uma segunda forma de type assertion que retorna
dois valores: o valor concreto (se a asserção for bem-sucedida) e um booleano que
indica se a asserção foi bem-sucedida ou não.
valor, ok := v.(T)

Se v contiver um valor do tipo T, valor receberá o valor concreto e ok será true.

Se v não contiver um valor do tipo T, a operação não causará pânico.
Em vez disso, valor receberá o valor zero do tipo T
(por exemplo, 0 para int, "" para string, nil para ponteiros) e ok será false.
*/
package main

import "fmt"

func main() {
	var i interface{} = "Olá mundo!"

	//Aserção bem sucedida
	s := i.(string)
	fmt.Println(s)

	//Tentativa com tipo incorreto, causará panico
	//f := i.(float64)
	//fmt.Println(f)

	f, ok := i.(float64)
	if ok {
		fmt.Printf("A asserção para float64 foi bem sucedida: %f\n", f)
	} else {
		fmt.Printf("A asserção para float64 falhou. Valor: %f\n", f)
	}

}
