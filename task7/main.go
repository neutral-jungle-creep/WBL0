package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex // мьютекс, который будет защищать элементы кеша
	items      map[int]string
}

func NewCache(i map[int]string) *Cache {
	return &Cache{
		items: i,
	}
}

func (c *Cache) Set(key int, value string) {
	c.Lock() // блок доступа к чтению и записи потому что с мапой уже работает одна из горутин
	c.items[key] = value
	c.Unlock() // разблокировка доступа для того, чтобы горутина, которая ждала, приступила к работе с мапой
}

func main() {
	cache := NewCache(map[int]string{})

	log.Println("writing starts")
	for i := 1; i < 4; i++ { // создание трех горутин
		go func(id int) {
			for j := 0; j < 70; j++ {
				// запись по рандомному ключу (поэтому некоторые значения могут перезаписаться)
				// значение - это номер воркера, чтобы было видно, какая из горутин записала
				cache.Set(rand.Intn(50000), " --- worker"+strconv.Itoa(id))
				// сразу лог о том, какой из воркеров пишет, чтобы было видно как реализуется конкурентность,
				//так как мапа неупорядоченная коллекция, то
				// в выводе ее элементов после записи, не будет понятно, что значения были записаны конкурентно
				log.Printf("worker [%d] write to cache", id)
			}

		}(i)
	}

	time.Sleep(5 * time.Second)

	log.Println("print cache items")
	for k, v := range cache.items {
		fmt.Println(k, ":", v) // вывод того, что получилось, чтобы убедиться, что значения были записаны

	}
}
