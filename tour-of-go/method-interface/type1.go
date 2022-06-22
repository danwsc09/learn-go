package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // "hello"

	s, ok := i.(string)
	fmt.Println(s, ok) // "hello", true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	var j interface{} = 99
	m, ok := j.(string)
	fmt.Printf("m: %q, ok: %v\n", m, ok)

	/*
		f = i.(float64) // panic, because i can't be float64
		fmt.Println(f)
	*/

	myMap := map[string]string{}
	myMap["foo"] = "bar"

	val, exists := myMap["foo"]
	fmt.Println("val, exists:", val, exists)

	val, exists = myMap["zzzzz"]
	fmt.Println("val, exists:", val, exists)

}
