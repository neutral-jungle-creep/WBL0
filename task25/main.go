package main

import (
	"context"
	"log"
	"time"
)

func sAfter(t time.Duration) {
	<-time.After(t) // когда время выйдет, в канал поступит сигнал и функция завершится
}

func sNewTimer(d time.Duration) {
	<-time.NewTimer(d).C // c - это канал таймера, по истечению таймера туда отправится сигнал
}

func formatTimeNow() string { // получение текущего времени и форматирование в функции, чтобы избежать дублирования
	return time.Now().Format("15:04:05")
}

func sWithTimeoutContext(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	<-ctx.Done() // когда время контекста истечет, сигнал поступил=т в этот канал
	cancel()     // и сработает функция
}

func sWithDeadLineContext(d time.Duration) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(d))
	<-ctx.Done() // когда время станет равно (время инициализации контекста) + (d)
	cancel()
}

func main() {
	log.Printf("func sAfter has been start at %s\n", formatTimeNow())
	sAfter(5 * time.Second)
	log.Printf("func sAfter finished at %s\n", formatTimeNow())

	log.Printf("func sNewTimer has been start at %s\n", formatTimeNow())
	sNewTimer(5 * time.Second)
	log.Printf("func sNewTimer finished at %s\n", formatTimeNow())

	log.Printf("func sWithDeadLineContext has been start at %s\n", formatTimeNow())
	sWithDeadLineContext(5 * time.Second)
	log.Printf("func sWithDeadLineContext finished at %s\n", formatTimeNow())

	log.Printf("func sWithTimeoutContext has been start at %s\n", formatTimeNow())
	sWithTimeoutContext(5 * time.Second)
	log.Printf("func sWithTimeoutContext finished at %s\n", formatTimeNow())
}
