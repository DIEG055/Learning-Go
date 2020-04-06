package main

import "fmt"

func main() {
	defer fmt.Println("World") // ejecutar este fragmento hasta que la funcion acabe
	fmt.Println("Hello")
}
