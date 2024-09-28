package main

import "fmt"

func main() {
	var people = map[string]int{
		"Dude": 47,
		"Chel": 100,
	}

	var name string
	var age int

	fmt.Print("Введите имя и возраст через пробел: ")
	fmt.Scan(&name, &age)

	people[name] = age

	for key, val := range people {
		fmt.Printf("Name: %s\tage: %d\n", key, val)
	}
}
