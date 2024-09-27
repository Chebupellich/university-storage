package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Printf("Число %d - четное", a)
	} else {
		fmt.Printf("Число %d - нечетное", a)
	}
}
