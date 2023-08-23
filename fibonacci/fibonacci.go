package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	cont, a, b := 0, 0, 1
	return func() int {
		if cont == 0 {
			cont++
			return a
		}
		if cont == 1 {
			cont++
			return b
		}
		c := a + b
		a, b = b, c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
