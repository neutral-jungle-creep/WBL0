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
	sync.Mutex
	items map[int]string
}

func NewCache(i map[int]string) *Cache {
	return &Cache{
		items: i,
	}
}

func (c *Cache) Set(key int, value string) {
	c.Lock()
	c.items[key] = value
	c.Unlock()
}

func main() {
	cache := NewCache(map[int]string{})

	log.Println("writing starts")
	for i := 1; i < 4; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				cache.Set(rand.Intn(50000), " --- worker"+strconv.Itoa(id))
				log.Printf("worker [%d] write to cache", id)
			}

		}(i)
	}

	time.Sleep(5 * time.Second)

	log.Println("print cache items")
	for k, v := range cache.items {
		fmt.Println(k, ":", v)

	}
}
