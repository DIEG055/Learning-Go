package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
// Un error es un valor que se puede describir a sí mismo usando una cadena de caracteres. Esta es la idea detrás del tipo error predefinido en Go. Es un interfaz con un sólo método que devuelve una cadena de caracteres:

// Las funciones de salida del paquete fmt llaman automáticamente a ese método cuando reciben un error.