package main

import "fmt"

func main() {
	var s string = "Hello ☀️"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("len of s:", len(s))

	var a rune = 'x'
	var ss string = string(a)
	var b byte = 'y'
	var ss2 string = string(b)
	fmt.Println("ss:", ss)
	fmt.Println("ss2:", ss2)

	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println("bs:", bs)
	fmt.Println("rs:", rs)

	var runeArr []rune = []rune{72, 101, 108, 108, 111, 44, 32, 127774}
	var runeArrStr string = string(runeArr)
	fmt.Println("rune str:", runeArrStr)

	var byteArr []byte = []byte{72, 101, 108, 108, 111, 44, 32, 240, 159, 140, 158}
	var byteArrStr string = string(byteArr)
	fmt.Println("byte str:", byteArrStr)
}
