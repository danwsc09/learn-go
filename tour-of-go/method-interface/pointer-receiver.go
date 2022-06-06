package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func (v Vertex) Scale(f float64) {
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
*/

func main() {
	/*
		v1 := Vertex{X: 5, Y: 12}
		fmt.Println(Abs(v1))
		Scale(&v1, 10)
		fmt.Println(Abs(v1))
	*/

	v1 := Vertex{X: 5, Y: 12}
	p1 := &Vertex{X: 3, Y: 4}
	v1.Scale(10)
	fmt.Println(v1.Abs())
	fmt.Println(v1)
	p1.Scale(10)
	fmt.Println(p1.Abs())
	fmt.Println(p1)

}
