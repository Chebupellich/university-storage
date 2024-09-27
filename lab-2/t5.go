package main

import (
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func (R Rectangle) Square() float64 {
	return R.width * R.height
}

func main() {
	var rect = Rectangle{width: 2, height: 2}
	fmt.Printf("Площадь - %.2f", rect.Square())
}
