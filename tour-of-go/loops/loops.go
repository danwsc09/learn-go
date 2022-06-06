package main

import (
	"fmt"
	"runtime"
	"time"
)

func Sqrt(x float64) float64 {
	cur := 1.0
	i := 0
	diff := 0.0000000001
	fmt.Println("Printf!")
	// for ; i < 10; i++ {
	// 	cur -= (cur*cur - x) / (2 * cur)
	// 	fmt.Println(cur)
	// }
	for cur*cur-x > diff || cur*cur-x < -diff {
		i++
		cur -= (cur*cur - x) / (2 * cur)
	}
	return cur
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away.")
	}
}

/*
	z -= (z * z - x) / (2 * z)
*/
