package main

import (
	"fmt"
	"math/rand"
)

func randGen(number int, ch chan int) {
	defer close(ch)
	for i := 0; i < number; i++ {
		ch <- rand.Intn(100)
	}
}

func isEven(value <-chan int, res chan<- string) {
	defer close(res)

	for val := range value {
		if val%2 == 0 {
			res <- fmt.Sprint("Число ", val, " - четное")
		} else {
			res <- fmt.Sprint("Число ", val, " - нечетное")
		}
	}
}

func main() {
	intCh := make(chan int)
	stringCh := make(chan string)

	go randGen(10, intCh)
	go isEven(intCh, stringCh)

	for {
		select {
		case num, ok := <-intCh:
			if ok {
				fmt.Println("Создано число:", num)
			} else {
				intCh = nil
			}
		case str, ok := <-stringCh:
			if ok {
				fmt.Println(str)
			} else {
				stringCh = nil
			}
		}

		if intCh == nil && stringCh == nil {
			break
		}
	}
}
