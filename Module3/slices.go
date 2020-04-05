package main

import "fmt"

func main() {
	p,pp := []int{2, 3, 5, 7, 11, 13} , []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",i, p[i])
	}


	fmt.Println("pp ==", pp)
	fmt.Println("pp[1:4] ==", pp[1:4])

	// missing low index implies 0
	fmt.Println("pp[:3] ==", pp[:3])

	// missing high index implies len(s)
	fmt.Println("pp[4:] ==", pp[4:])
}