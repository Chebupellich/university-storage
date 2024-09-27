package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a []int
	for i := 0; i < 5; i++ {
		a = append(a, rand.Intn(100))
	}

	fmt.Println(a)

	a = append(a[:3], a[3+1:]...)
	fmt.Println(a)
}
