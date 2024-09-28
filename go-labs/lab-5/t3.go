package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var circle = Circle{radius: 5.0}
	fmt.Printf("Радиус: %.2f\nПлощадь: %.2f", circle.radius, circle.Area())
}
