package main

import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	p := Vertex{1, 2}
// 	q := &p
// 	// fmt.Println(q)
// 	q.X = 1e9
// 	fmt.Println(p)
// }

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertex
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}