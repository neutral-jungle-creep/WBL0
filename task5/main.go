package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func sendRandomSymbolToChan(c context.Context, ch chan<- int) {
	defer close(ch) // при завершении горутины канал гарантированно закроется

	for {
		select {
		case <-c.Done(): // если пришло сообщение о завершении контекста, функция вернет ошибку контекста
			log.Printf("time is up: %s", c.Err().Error())
			return
		default:
			ch <- 1 + rand.Intn(3) // отправка рандомного числа в канал от 1 до 3
		}
	}
}

func main() {
	var workTime time.Duration                     // объявление переменной для времени работы программы
	if _, err := fmt.Scan(&workTime); err != nil { // получение длительности работы программы из консоли от пользователя
		log.Fatal(err)
	}

	ch := make(chan int) // инициализация канала
	ctx, cancel := context.WithTimeout(context.Background(), workTime*time.Second)

	go sendRandomSymbolToChan(ctx, ch)

	for d := range ch {
		fmt.Printf("program prints --- %v\n", string(rune(d))) // печать на консоль случайного символа
	}
	cancel()
}
