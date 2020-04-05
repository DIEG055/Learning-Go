package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // el while de C se deletrea for en Go.
		sum += sum
	}
	fmt.Println(sum)

	// for {   Bucle infinito
	// }
}	