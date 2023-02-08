package main

import "fmt"

func main() {
	// инициализация массива с числами и переменной типа int, в которой будет храниться результат
	var numbers = [5]int{2, 4, 6, 8, 10}
	var result int

	// инициализация канала, работающего с типом int
	var ch = make(chan int, len(numbers))

	for _, i := range numbers {
		go func(num int) {
			// запись результата вычисления в канал
			ch <- num * num
		}(i)
	}

	for i := 0; i < len(numbers); i++ {
		result = result + <-ch
	}

	fmt.Println(result)
}
