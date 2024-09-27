package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}
