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

// создание пользовательского типа worker
// воркеру каналы нужны только для чтения, поэтому преобразовала канал в однонаправленный
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
			fmt.Printf("[%d] worker prints --- digit %d\n", w.id, num) // чтение числа из канала и вывод на консоль
		}
	}
}

func sendNumsToChan(osSignalCh <-chan os.Signal, numsCh chan<- int) {
	for {
		select {
		case <-osSignalCh: // если пришел сигнал в этот канал, то программа выйдет из бесконечного цикла
			return
		default:
			numsCh <- rand.Intn(500) // отправка данных в канал, который читают горутины
		}
	}
}

func main() {
	var workers int                               // объявление переменной для кол-ва воркеров
	if _, err := fmt.Scan(&workers); err != nil { // получение числа воркеров из консоли от пользователя
		log.Fatal(err)
	}

	osSignalCh := make(chan os.Signal, 1)                      // инициализация канала для сигнала ОС
	signal.Notify(osSignalCh, syscall.SIGINT, syscall.SIGTERM) // отправка сигнала ОС в канал

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

	//todo дописать комменты
	go func() {
		<-osSignalCh
		close(osSignalCh)
	}()

	sendNumsToChan(osSignalCh, numsCh)
	fmt.Println("Break main loop")

	wg.Wait() // ожидание завершения работы всех воркеров
}
