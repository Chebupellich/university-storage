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
	sum := 0

	for _, val := range numbers {
		num, err := strconv.Atoi(val)
		if err == nil {
			sum += num
		} else {
			panic("Некорректный ввод")
		}
	}

	fmt.Print("Сумма элементов: ", sum)
}
