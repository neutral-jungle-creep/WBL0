package main

import (
	"fmt"
	"log"
	"math/rand"
)

func makeNumsMap(set1, set2 [10]int) map[int]int {
	var result = map[int]int{} // инициализация мапы для промежуточного результата

	for _, item := range set1 {
		result[item] = 0 // добавление элемента, где ключ - значение из массива1, а значение 1
	}

	for _, item := range set2 {
		if _, exist := result[item]; exist { // если ключ уже есть в мапе
			result[item] = 1 // изменить его значение на 1, таким образом у всех пересечений будет значение 1
		}
	}
	return result
}

func main() {
	var set1, set2 = [10]int{}, [10]int{} // инициализация двух массивов размером 10 элементов

	for i := 0; i < 10; i++ {
		set1[i] = rand.Intn(5)
		set2[i] = rand.Intn(15)
	}
	log.Printf("set1 = %v", set1)
	log.Printf("set2 = %v", set2)

	for i, item := range makeNumsMap(set1, set2) {
		if item == 1 { // вывод на консоль результата пересечения двух множеств
			fmt.Println(i)
		}
	}
}
