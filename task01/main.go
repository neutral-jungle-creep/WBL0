package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       uint8
}

type Action struct {
	// Human - тип анонимного поля, внедренного в структуру
	Human
}

// run метод для типа Human
func (h *Human) run() {
	fmt.Printf("%s is running now", h.firstName)
}

func main() {
	var action = Action{
		Human: Human{
			firstName: "Артур",
			lastName:  "Григорьев",
			age:       23,
		},
	}

	action.run() // вызов метода анонимного поля Human
}
