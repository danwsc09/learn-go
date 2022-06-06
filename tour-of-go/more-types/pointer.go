package main

import "fmt"

func main() {
	var i, j int = 10, 100
	fmt.Println("i, j:", i, j)
	p := &i
	fmt.Println("pointer p:", p)

	*p = *p + 999
	fmt.Printf("at %p: %d\n", p, *p)
	fmt.Println("i:", i)

	p = &j
	*p = 1
	fmt.Printf("*p: %d, j: %d\n", *p, j)
}
