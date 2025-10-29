package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Resposta"] = 42
	fmt.Println("O valor da resposta:", m["Resposta"])

	m["Resposta"] = 48
	fmt.Println("O valor da resposta:", m["Resposta"])

	delete(m, "Resposta")
	fmt.Println("O valor da resposta:", m["Resposta"])

	//Se key está em m, ok é true. Se não, ok é false.
	//Se key não está no map então elem tem valor zero para o elemento do tipo do map.
	v, ok := m["Resposta"]
	fmt.Println("O valor da resposta:", v, "Presente?", ok)
}
