package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

// создала структуру воркера чтобы не передавать в функцию service 4+ аргументов
type worker struct {
	id          int
	numsQueueCh chan int
	shoutDownCh chan os.Signal
	wg          *sync.WaitGroup
}

// newWorker функция-конструктор, которая гарантированно вернет ссылку на объект типа worker
func newWorker(id int, numsCh chan int, osSignalCh chan os.Signal, wg *sync.WaitGroup) *worker {
	return &worker{
		id:          id,
		numsQueueCh: numsCh,
		shoutDownCh: osSignalCh,
		wg:          wg,
	}
}

// метод структуры worker читает канал, в которые посылаются случайные int и печатает логи в консоль
func (w *worker) service() {
	defer w.wg.Done() // при завершении работы воркера счетчик группы ожидания гарантированно уменьшится на 1

	fmt.Printf("worker [%d] is starting!\n", w.id)

	for {
		select {
		case <-w.shoutDownCh: // если пришел сигнал в этот канал, то горутина будет завершена
			fmt.Printf("shutting down [%d] worker!\n", w.id)
			return
		case num := <-w.numsQueueCh:
			fmt.Printf("[%d] worker prints --- digit %s\n", w.id, string(rune(num)))
		}
	}
}

func sendRandomNumsToChan(osSignalCh <-chan os.Signal, numsCh chan<- int) {
	for {
		select {
		case <-osSignalCh: // если пришел сигнал в этот канал, то программа выйдет из бесконечного цикла
			return
		default:
			numsCh <- 1 + rand.Intn(3) // отправка данных в канал, который читают горутины
		}
	}
}

func main() {
	var workers int
	if _, err := fmt.Scan(&workers); err != nil { // получение числа воркеров из консоли от пользователя
		log.Fatal(err)
	}

	osSignalCh := make(chan os.Signal, 1) // инициализация канала для сигнала ОС
	// передача функции канала, из которого она будет ждать сигнал для завершения работы
	signal.Notify(osSignalCh, syscall.SIGINT, syscall.SIGTERM)

	// инициализация канала для работы с типом int (генерируемых программой данных)
	numsCh := make(chan int)

	// инициализация группы ожидания
	var wg = &sync.WaitGroup{}

	// создание введенного пользователем числа воркеров
	// счет решила начать с единицы, потому что так привычнее пользователю
	for i := 1; i <= workers; i++ {
		wg.Add(1) // увеличение счетчика у группы ожидания

		wrk := newWorker(i, numsCh, osSignalCh, wg) // создание переменной типа worker с помощью конструктора
		go wrk.service()
	}

	go sendRandomNumsToChan(osSignalCh, numsCh)

	// здесь основной поток остановится и будет ждать сигнала от ОС
	<-osSignalCh
	close(osSignalCh)

	fmt.Println("program was finish")

	wg.Wait() // ожидание завершения работы всех воркеров
}
