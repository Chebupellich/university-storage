package main

import (
	"fmt"
)

func Mean(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	a := 4
	b := 20

	fmt.Printf("a: %d, b: %d\nMean value: %.2f", a, b, Mean(a, b))
}
