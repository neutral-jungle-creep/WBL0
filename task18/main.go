package main

import (
	"fmt"
	"sync"
)

type counter struct { // инкапсулирование структуры счетчика для того, из других пакетов счетчик нельзя было создать без интерфейса
	mx    sync.Mutex
	value int32
}

type Counter interface { // создание интерфейса с доступными из вне функциями
	GetCounterValue() int32
	Increment()
}

func NewCounter() Counter { // конструктор интерфейса, который инициализирует структуру counter
	return &counter{
		value: 0,
	}
}

func (c *counter) GetCounterValue() int32 {
	return c.value
}

func (c *counter) Increment() {
	c.mx.Lock()
	c.value += 1
	c.mx.Unlock()
}

func testingCounter(wg *sync.WaitGroup, count Counter) {
	for i := 0; i < 50; i++ {
		count.Increment()
	}
	wg.Done()
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

func main() {
	count := NewCounter()

	var wg = &sync.WaitGroup{}

	for i := 0; i < 10; i++ { // запуск 10 потоков
		wg.Add(1)
		go testingCounter(wg, count)
	}

	wg.Wait()

	fmt.Printf("counter's value = [%d]", count.GetCounterValue())
}
