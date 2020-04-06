package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { //recover useful in functions
			log.Println("Error: ", err)
			// panic(err)
		}
	}()
}

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")

}
