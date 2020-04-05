package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {S
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
// Los slices se crean con la función make. Ésta reserva memoria para una tabla inicializada a cero y devuelve un slice que apunta a esa tabla:

// Los slices tienen una longitud y una capacidad. La capacidad es el tamaño máximo que puede abarcar el slice en la tabla subyacente.