package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	//f := MyFloat(-3.2)
	v := MyVertex{3, 4}

	a = &v
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyVertex struct {
	X, Y float64
}

func (v *MyVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
