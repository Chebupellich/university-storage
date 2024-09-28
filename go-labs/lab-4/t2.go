package main

import "fmt"

func MapMean(people map[string]int) float64 {
	if len(people) == 0 {
		return -1
	}
	acc := 0
	for _, val := range people {
		acc += val
	}
	return float64(acc) / float64(len(people))
}

func main() {
	var people = map[string]int{
		"Dude":  47,
		"Chel":  100,
		"Bloba": 11,
	}

	for key, val := range people {
		fmt.Printf("Name: %s\tage: %d\n", key, val)
	}
	fmt.Printf("Среднее по больнице: %.2f", MapMean(people))
}
