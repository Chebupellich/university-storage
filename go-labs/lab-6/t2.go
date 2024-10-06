package main

import "fmt"

func rfactorial(a int, ch chan int) {
	defer close(ch)

	acc := 1
	for i := 1; i <= a; i++ {
		acc *= i
		ch <- acc
	}
}

func main() {
	intCh := make(chan int)

	go rfactorial(5, intCh)

	for num := range intCh {
		fmt.Println(num)
	}
}
