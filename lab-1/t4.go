package main

import (
	"fmt"
)

func main() {
	num1, num2 := 3, -10

	fmt.Println("Сложение: ", num1+num2)
	fmt.Println("Вычитание: ", num1-num2)
	fmt.Println("Умножение: ", num1*num2)
	fmt.Println("Деление: ", num2/num1)
	fmt.Println("Деление с остатком: ", num1%num2)
}
