/*
Um map mapeia chaves para valores.
O valor zero de um map é nil. Um map nil não tem chaves, nem chaves podem ser adicionadas.
A função make retorna um map com um tipo determinado, inicializado e pronto para o uso.
*/
package main

import "fmt"

type Vortex struct {
	Lat, Long float64
}

var m map[string]Vortex

func main() {
	m = make(map[string]Vortex)
	m["Bell Labs"] = Vortex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}
