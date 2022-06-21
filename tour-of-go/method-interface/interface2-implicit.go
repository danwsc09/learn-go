package main

import "fmt"

// implicit interface
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

// interface value
type J interface {
	N()
}

type U struct {
	S string
}

type F float64

func (u *U) N() {
	fmt.Println(u.S)
}

func (f F) N() {
	fmt.Println(f)
}

func describe(j J) {
	fmt.Printf("value, type: (%v, %T)\n", j, j)
}

func main() {
	// T implicitly implements I
	var i I = T{S: "HELLO!!!"}
	i.M()

	// interface values
	var j J = &U{S: "WORLD!!!"}
	describe(j)
	j.N()

	j = F(-2.3)
	describe(j)
	j.N()
}
