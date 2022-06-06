package main

import (
	"fmt"
)

func main() {
	a := "some string"
	fmt.Println(len(a))
	b := make(map[string]int)
	if b["hi"] == 0 {
		fmt.Println("b is nil")
	}
}
