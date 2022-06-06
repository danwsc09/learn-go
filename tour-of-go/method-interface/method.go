package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I am %s and %d years old.", p.name, p.age)
}

func Introduce(p Person) string {
	return fmt.Sprintf("Hi, I am %s and %d years old.", p.name, p.age)
}

type CustomFloat float64

func (c CustomFloat) Abs() float64 {
	if c < 0 {
		return float64(-c)
	}
	return float64(c)
}

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

func main() {
	p1 := Person{name: "Dan", age: 30}
	p2 := Person{"Phil", 33}
	fmt.Println(p1.Introduce())
	fmt.Println(p2.Introduce())
	fmt.Println(Introduce(p1))
	fmt.Println(Introduce(p2))

	f := CustomFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
