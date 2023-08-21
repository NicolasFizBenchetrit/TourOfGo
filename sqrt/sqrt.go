package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	prev := 0.0
	i := 0
	for ; !(math.Abs(z-prev) < 1e-08); i++ {
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println("Iterations ->", i)
	return z
}

func main() {
	fmt.Println("Mine -> ", Sqrt(10))
	fmt.Println("Library math -> ", math.Sqrt(10))
}
