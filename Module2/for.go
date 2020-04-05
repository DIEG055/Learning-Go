package main

import "fmt"

func main() {
	ss,sum := 1,2
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for ; ss < 10; { // Como en C o Java, puedes dejar vacías las sentencias pre y post.
		ss += ss
	}
	fmt.Println(ss)
}