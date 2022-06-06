package main

import (
	"fmt"
)

func main() {
	var a [5]string
	fmt.Println(a)
	a[3] = "99"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{"dan", "matt", "james", "john"}
	fmt.Println(names)

	f2 := names[0:2]
	f3 := names[2:4]
	fmt.Println(f2, f3)

	f2[0] = "changed!"
	f3[0] = "this too!"
	// [changed, matt, this too, john]
	fmt.Println(f2, f3)
	fmt.Println("names:", names)

	arr1 := []int{2, 3, 5, 7, 11, 13}
	arr1 = arr1[1:4]
	fmt.Println(arr1)
	arr1 = arr1[:2]
	fmt.Println(arr1)
}
