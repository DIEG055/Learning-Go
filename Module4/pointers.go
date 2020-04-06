package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	var ms *myStruct
	fmt.Println(ms)    // nil
	ms = new(myStruct) // new crea un apuntator a la estructura
	fmt.Println(ms)
	(*ms).foo = 50 // tambien se puede sin  (* )  -> dereference the pointer
	fmt.Println((*ms).foo)
	ms.foo = 555 // tambien se puede sin  (* )
	fmt.Println(ms.foo)

}

type myStruct struct {
	foo int
}
