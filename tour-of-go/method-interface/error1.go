package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// error type(built-in interface) implements: Error() string;
// similar to Stringer interface's : String() string
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didnt work!!"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	var e error
	myE := &MyError{time.Now(), "it didnt workerino :("}
	e = myE
	fmt.Println(e)

	i, err := strconv.Atoi("ab")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}
