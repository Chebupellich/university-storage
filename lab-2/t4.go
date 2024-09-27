package main

import (
	"fmt"
)

func main() {
	var a string

	fmt.Print("Введите строку: ")
	fmt.Scan(&a)

	fmt.Print("Длина строки - ", len(a))
}
