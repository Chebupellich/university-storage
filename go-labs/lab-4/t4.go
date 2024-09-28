package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	fmt.Print("Измененная строка: ", strings.ToUpper(input))
}
