package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func factorial(a int, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	acc := 1
	for i := 1; i < a; i++ {
		acc *= i
	}
	res <- fmt.Sprint("Factorial of ", a, ": ", acc)
}

func random(res chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	res <- fmt.Sprint("Rand number: ", rand.Intn(100))
}

func sum(nums []int, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	acc := 0
	for i := 0; i < len(nums); i++ {
		acc += nums[i]
	}
	res <- fmt.Sprint("Sum: ", acc)
}

func main() {
	var wg sync.WaitGroup
	stringChannel := make(chan string, 3)

	wg.Add(3)

	go factorial(5, stringChannel, &wg)
	go random(stringChannel, &wg)
	go sum([]int{1, 2, 3, 4, 5}, stringChannel, &wg)
	wg.Wait()

	close(stringChannel)
	for val := range stringChannel {
		fmt.Println(val)
	}
}
