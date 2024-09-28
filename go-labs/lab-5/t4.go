package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	a float64
	b float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func main() {
	var shapes = []Shape{Circle{radius: 2}, Rectangle{a: 3, b: 4}, Square{side: 5}}
	for _, val := range shapes {
		fmt.Printf("Площадь: %.2f\n", val.Area())
	}
}
