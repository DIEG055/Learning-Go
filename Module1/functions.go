package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(flag bool,x, y int) int { // Cuando dos parámetros o más tienen el mismo tipo, puedes omitirlo en todos excepto el último.
	return x + y + y
}


func main() {
	fmt.Println(add2(true,42, 13))
}