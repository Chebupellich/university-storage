package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getInfo() string {
	return fmt.Sprintf("Имя: %s\tВозраст: %d", p.name, p.age)
}

func main() {
	var dude = Person{name: "Dude", age: 22}
	fmt.Print(dude.getInfo())
}
