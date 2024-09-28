package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Имя: %s\tВозраст: %d", p.name, p.age)
}

func (p *Person) Birthday() {
	(*p).age++
}

func main() {
	var dude = Person{name: "Dude", age: 22}
	fmt.Println("--- ДО ---")
	fmt.Print(dude.getInfo())

	dude.Birthday()

	fmt.Println("\n--- ПОСЛЕ ---")
	fmt.Print(dude.getInfo())
}
