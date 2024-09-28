package main

import "fmt"

func main() {
	var people = map[string]int{
		"Dude":  47,
		"Chel":  100,
		"Bloba": 11,
	}

	var name string

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	delete(people, name)

	for key, val := range people {
		fmt.Printf("Name: %s\tage: %d\n", key, val)
	}
}
