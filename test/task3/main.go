package main

import "fmt"

func main() {
	var numbers = [5]int{2, 4, 6, 8, 10} // инициализация массива с числами и
	var result int                       // объявление переменной типа int, в которой будет храниться результат

	ch := make(chan int, len(numbers)) // инициализация канала, работающего с типом int

	for _, i := range numbers {
		go func(num int) {
			ch <- num * num // запись результата вычисления в канал
		}(i)
	}

	for i := 0; i < len(numbers); i++ {
		result = result + <-ch // суммирование результата
	}

	fmt.Println(result)
}
