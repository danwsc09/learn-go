package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Stringer interface {
	String() string
}

func (s Student) String() string {
	return fmt.Sprintf("%v (%v years)", s.Name, s.Age)
}

func main() {
	var s Stringer
	a := Student{
		Name: "Pew",
		Age:  30,
	}
	b := Student{
		Name: "Pa",
		Age:  20,
	}
	fmt.Println(a, b)
	s = a
	fmt.Println(s)
	s = b
	fmt.Println(s)
}
