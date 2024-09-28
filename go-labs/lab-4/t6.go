package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите целые числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")
	intNumbers := make([]int, len(numbers))

	for index, val := range numbers {
		num, err := strconv.Atoi(val)
		if err == nil {
			intNumbers[index] = num
		} else {
			panic("Некорректный ввод")
		}
	}

	for i, j := 0, len(intNumbers)-1; i < j; i, j = i+1, j-1 {
		intNumbers[i], intNumbers[j] = intNumbers[j], intNumbers[i]
	}

	fmt.Print("Числа в обратном порядке: ", intNumbers)
}
