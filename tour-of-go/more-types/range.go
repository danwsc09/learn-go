package main

import (
	"fmt"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow {
		fmt.Printf("v: %d\n", v)
		// fmt.Printf("i: %d, v: %d\n", i, _)
		// fmt.Printf("2 ** %d = %d\n", i, v)
	}
}
