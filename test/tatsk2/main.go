package main

import "fmt"

func main() {
	// объявление массива с числами
	var numbers = [5]int{2, 4, 6, 8, 10}

	// инициализация канала типа int
	ch := make(chan int)

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
}
