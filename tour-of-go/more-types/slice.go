package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	var what []int
	fmt.Println(what, len(what), cap(what))
	if what == nil {
		fmt.Println("what is nil")
	}

	fmt.Println("\n--------- MAKE ---------")
	fmt.Println()
	// make a slice of length 5
	a := make([]int, 5)
	fmt.Printf("cap: %d, len: %d\n", cap(a), len(a))

	// make a 2d array of size dy, dx
	dy, dx := 10, 8
	b := make([][]int, dy)
	for y := 0; y < dy; y++ {
		arr := make([]int, dx)
		b[y] = arr
	}
	fmt.Println(b)
}
