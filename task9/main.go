package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	numbers := make([]int, 0, 22) // инициализация пустого слайса с емкостью 22

	for i := 0; i < 22; i++ {
		numbers = append(numbers, rand.Intn(100)) // заполнение слайса рандомными числами от 0 до 100
	}

	chReadNum := make(chan int) // инициализация двух каналов
	chPrintNum := make(chan int)

	go func() {
		for n := range chReadNum { // получение данных по одному из канала
			log.Printf("chReadNum gets num --- %d\n", n) // логиговарие полученного
			newNum := n * 2                              // умножение
			chPrintNum <- newNum                         // отправка в канал для печати на консоль
		}
	}()
	go func() {
		for n := range chPrintNum {
			fmt.Printf("new number from chPrintNum = [%d]\n", n) // печать на консоль результата умножения
		}
	}()

	for _, num := range numbers { // отправка чисел из массива в канал
		chReadNum <- num
		time.Sleep(1 * time.Second)
	}

}
