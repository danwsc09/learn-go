package main

import "fmt"

func main() {
	defer fmt.Println("finally done with main")
	fmt.Println("Hi")
	// f()
}

func f() {
	defer fmt.Println("defer in f")
	fmt.Println("Funsies")

	fmt.Println("Will throw an error")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover!", r)
		}
	}()
	panic("Im panicking!")
	fmt.Println("does this run?")
}
