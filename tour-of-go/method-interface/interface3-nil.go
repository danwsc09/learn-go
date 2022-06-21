package main

import "fmt"

type MyI interface {
	M()
}

type MyT struct {
	S string
}

func (t *MyT) M() {
	if t == nil {
		fmt.Println("<NIL>")
		return
	}
	fmt.Println(t.S)
}

func describeI(i MyI) {
	fmt.Printf("value, type: (%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("value, type: (%v, %T)\n", i, i)
}

func main() {
	var i MyI
	// RUNTIME ERROR - no concrete M() to call
	//describeI(i)
	//i.M()

	var t *MyT
	i = t
	describeI(i)
	i.M()

	// an interface that holds a nil concrete value is non-nil !
	if i == nil {
		fmt.Println("i is NILLLLL")
	}

	i = &MyT{"hello!!"}
	describeI(i)
	i.M()

	// empty interface
	var IF interface{}
	describeEmpty(IF)

	IF = 42
	describeEmpty(IF)

	IF = "HELLO THEREEE"
	describeEmpty(IF)

}
