package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	ver := Vertex{x: 1, y: 2}
	ver.x = 10
	fmt.Println("ver:", ver)
	p := &ver
	p.y = 10
	fmt.Println("ver:", ver)

	fmt.Println(v1, v2, v3, p)
}
