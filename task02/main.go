package main

import "fmt"

func main() {
	var numbers = [5]int{2, 4, 6, 8, 10} // инициализация массива с числами
	ch := make(chan int, len(numbers))   // инициализация канала, работающего с типом int

	for _, i := range numbers {
		go func(num int) {
			ch <- num * num // запись результата вычисления в канал
		}(i)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-ch) // чтение и вывод данных из канала
	}
	close(ch) // закрытие канала
}
