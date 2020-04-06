package main

import "fmt"

func main() {
	 students := [...]int {1,2,3,4,5} /// crear un arrray del tamano de la cantidad de datos	
	 s2 := students // crea una copia, no apunta al original
	fmt.Print(students)
}