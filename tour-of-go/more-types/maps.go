package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	// var m map[string]string
	// m = make(map[string]string)

	var m = map[string]Vertex{
		"Bell labs": {1.2, 3.4},
		"Google":    Vertex{5.6, 7.8},
	}
	// m["hi"] = "pewpew"
	// m["hello"] = "byebe"

	m["pewpew"] = Vertex{333.3, 222.2}

	fmt.Println(m)

	myMap := make(map[string]string)
	myMap["what"] = "C++"
	myMap["how"] = "golang"
	myMap["mcr"] = "when i was a young boy"
	fmt.Println("myMap:", myMap)

	elem, ok := myMap["how"]
	fmt.Printf("The value: %s, ok: %s\n", elem, ok)
	fmt.Println("mymap[\"random\"]:", myMap["this doesntexist"])
	// fmt.Println("OK:", ok)
}
