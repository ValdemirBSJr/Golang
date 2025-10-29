package main

import "fmt"

// ChavesDoMapa extrai todas as chaves de um mapa para uma fatia.
//
// A declaração de tipo aqui é [K comparable, V any].
// - K comparable: O tipo da CHAVE (K) deve ser 'comparable', pois as chaves
//   de um mapa em Go sempre precisam ser comparáveis.
// - V any: O tipo do VALOR (V) pode ser 'any' (qualquer tipo), pois
//   os valores de um mapa não têm essa restrição. 'any' é um alias para 'interface{}'.
func ChavesDoMapa[K comparable, V any](mapa map[K]V) []K {
	// Cria uma fatia para armazenar as chaves, do tipo K.
	chaves := make([]K, 0, len(mapa))

	// Itera sobre o mapa. Em um 'range' de mapa, o primeiro valor é a chave.
	for K := range mapa {
		chaves = append(chaves, K)
	}
	return chaves
}

func main() {
	// Exemplo com mapa de string para int
	mapaDeContagem := map[string]int{
		"maçãs":    5,
		"bananas":  10,
		"laranjas": 8,
	}

	chavesStr := ChavesDoMapa(mapaDeContagem)
	fmt.Println("Chaves do mapa de contagem:", chavesStr)

	// A mesma função funciona perfeitamente com um mapa de int para bool
	mapaDeStatus := map[int]bool{
		200: true,
		404: false,
		503: false,
	}

	chavesInt := ChavesDoMapa(mapaDeStatus)
	fmt.Println("Chaves do mapa de status:", chavesInt)

}
