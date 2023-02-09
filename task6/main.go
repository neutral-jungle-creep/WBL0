package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func contextDoneRoutineStop() {
	var workTime = time.Duration(5) // явное назначение типа, иначе был бы int
	ch := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), workTime*time.Second)

	go func(c context.Context, ch chan int) {
		defer close(ch) // при завершении горутины канал гарантированно закроется

		for {
			select {
			case <-c.Done():
				log.Println("finish func contextDoneRoutineStop")
				return
			default:
				ch <- 1 + rand.Intn(3) // отправка рандомного числа в канал от 1 до 3
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx, ch)

	for d := range ch { // если не закрыть канал в горутине, то будет all goroutines are asleep - deadlock!
		fmt.Printf("func contextDoneRoutineStop prints --- %s\n", string(rune(d))) // печать на консоль случайного символа
	}

	cancel()
}

func sendDigitToChan(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- 1 + rand.Intn(3) // отправка рандомного числа в канал от 1 до 3
		time.Sleep(1 * time.Second)
	}
}

func closeChanRoutineStop() {
	ch := make(chan int, 5) // чтобы избежать  all goroutines are asleep - deadlock! нужно указать размер буфера

	go func() {
		for {
			num, ok := <-ch
			if !ok {
				log.Println("finish func cancelCloseChan")
				return
			}
			fmt.Printf("func cancelCloseChan prints --- %s\n", string(rune(num)))
		}
	}()

	sendDigitToChan(ch)

	close(ch)
	time.Sleep(1 * time.Second)
}

// TODO - прочитать про другие способы
func main() {
	contextDoneRoutineStop()
	closeChanRoutineStop()
}
