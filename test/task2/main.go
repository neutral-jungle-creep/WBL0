package main

import "fmt"

func main() {
	// инициализация массива с числами
	var numbers = [5]int{2, 4, 6, 8, 10}
	// инициализация канала, работающего с типом int
	var ch = make(chan int, 5)

	for _, i := range numbers {
		go func(num int) {
			// запись результата вычисления в канал
			ch <- num * num
		}(i)
	}

	for i := 0; i < len(numbers); i++ {
		// чтение и вывод данных из канала
		fmt.Println(<-ch)
	}
	// закрытие канала
	close(ch)
}
