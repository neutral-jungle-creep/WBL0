package main

import (
	"fmt"
)

func printType(i any) { // any является псевдонимом для interface{} и эквивалентен interface{} во всех смыслах (из официальной доки)
	fmt.Printf("%T\n", i) // %T - синтаксис вывода типа переменной
}

func main() {
	ch := make(chan float64)

	printType(5) // когда значение передается в эту функцию оно становится interface{}
	printType("test")
	printType(true)
	printType(ch)
}
