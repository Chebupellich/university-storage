package main

import (
	"fmt"

	"t2.go/mathutils"
	"t2.go/stringutils"
)

func main() {
	var a int
	var b string

	fmt.Print("Введите целое число: ")
	fmt.Scan(&a)

	fmt.Println("Факториал: ", mathutils.Factorial(a))

	fmt.Print("Введите строку: ")
	fmt.Scan(&b)

	fmt.Println("Перевернутая строка: ", stringutils.ReverseString(b))
}
