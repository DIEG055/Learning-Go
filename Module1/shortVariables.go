package main

import "fmt"

var vv =45

func main() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java,vv)
}

// Dentro de una función, la asignación especial := se puede utilizar para ahorrarse la declaración.

// (Fuera de las funciones, todas las sentencias empiezan con una palabra reservada y la asignación especial := no existe.)