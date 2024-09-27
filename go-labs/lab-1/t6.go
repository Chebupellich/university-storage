package main

import (
	"fmt"
)

func Mean(a, b, c int) float64 {
	return float64(a+b+c) / 3.0
}

func main() {
	a := 4
	b := 20
	c := -7

	fmt.Printf("a: %d, b: %d, c: %d\nMean value: %.2f", a, b, c, Mean(a, b, c))
}
