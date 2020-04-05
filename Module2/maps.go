package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var i map[int] float32

var mm = map[string]Vertex{
	"BellLabs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Vertex) // Antes de usar un map, hay que crearlo con make (en vez de con new); y el map nil por definición está vacío y no se le puede añadir nada.
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	i = make(map[int]float32)
	i[1] = 2.5
	fmt.Println(m["Bell Labs"])
	fmt.Println(i[1])
	

	fmt.Println(mm)
}