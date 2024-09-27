package main

import (
	"fmt"
)

func SumSub(a float32, b float32) [2]float32 {
	return [2]float32{a + b, a - b}
}

func main() {
	fmt.Println("Деление с остатком: ", SumSub(1.37, 0.63))
}
