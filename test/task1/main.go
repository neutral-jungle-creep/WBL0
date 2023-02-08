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
	// объявление переменной типа Human
	var artur = Human{
		firstName: "Артур",
		lastName:  "Григорьев",
		age:       23,
	}

	// объявление переменной типа Action
	var action = Action{
		Human: artur,
	}

	// вызов метода анонимного поля Human
	action.run()

}
