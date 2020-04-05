package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

// Go es un lenguaje compilado, por lo tanto tenemos que usar éste para generar un ejecutable del archivo.

// En la terminal exploraremos los siguientes comandos:

// $ go build [filename]
// $ go build main.go
// Este comando sirve para compilar el archivo que le indiquemos. Eso genera un ejecutable, es decir, un binario.

// $ go run [filename]
// $ go run main.go
// Este comando, al igual que el go build, también compila el archivo que le indiquemos con la diferencia de que también lo va a ejecutar. Resulta útil para probar rápidamente código en go. La principal diferencia, es que este comando compila el archivo, ejecuta el binario y elimina el archivo de manera inmediata.