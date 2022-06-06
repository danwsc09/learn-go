package main

import (
	"fmt"
)

func SomeFun() {
	defer fmt.Println("First")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("pewpew")
}

func main() {
	SomeFun()
	defer fmt.Println("World")

	fmt.Println("Hello")
}

// pewpew third second first hello world
