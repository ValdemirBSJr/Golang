// Maps literais são como structs literais, mas as chaves são obrigatórias.
package main

import "fmt"

type Vortex struct {
	Lat, Long float64
}

var m = map[string]Vortex{
	"Bell Labs": Vortex{
		40.68433, -74.39967,
	},
	"Google": Vortex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["Google"])
	fmt.Println(m["Bell Labs"].Long)
}
