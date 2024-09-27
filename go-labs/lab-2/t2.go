package main

import (
	"fmt"
)

func GetSign(a int) string {
	if a < 0 {
		return "Negative"
	} else if a > 0 {
		return "Positive"
	} else {
		return "Zero"
	}
}

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	fmt.Print("Знак - ", GetSign(a))
}
