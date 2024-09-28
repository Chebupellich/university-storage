package main

import "fmt"

type Stringer interface {
	ShowInfo()
}

type Book struct {
	name   string
	year   int
	author string
	text   string
}

func (b Book) ShowInfo() {
	fmt.Printf("------\nНазвание: %s\nГод: %dг\nАвтор: %s\nТекст: %s\n",
		b.name, b.year, b.author, b.text)
}

func main() {
	var strings = []Stringer{
		Book{
			name:   "Приключения Абобы",
			year:   1337,
			author: "Абоба",
			text:   "Жил был абоба и забыл",
		},
		Book{
			name:   "Сказка о забытом курсаче",
			year:   2024,
			author: "Все",
			text:   "Емае, уже курсач надо сдавать",
		},
	}

	for _, val := range strings {
		val.ShowInfo()
	}
}
