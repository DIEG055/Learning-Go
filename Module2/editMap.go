package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // Si la llave key está en m, ok será true, Si no, ok será false y elem será el valor cero del tipo de los elementos.
	fmt.Println("The value:", v, "Present?", ok)
}