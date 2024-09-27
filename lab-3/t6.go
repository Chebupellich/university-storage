package main

import (
	"fmt"
)

func main() {
	var a = []string{"Aboba", "NOT-aboba", "Hi-Victoria"}
	fmt.Println(a)

	maxIndex := 0
	for i := 0; i < len(a)-1; i++ {
		if len(a[i]) < len(a[i+1]) {
			maxIndex = i + 1
		}
	}

	fmt.Print("Самая длинная строка: ", a[maxIndex])
}
