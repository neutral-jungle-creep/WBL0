package main

import "fmt"

func main() {
	var (
		words    = [...]string{"cat", "cat", "dog", "cat", "tree"} // инициализация массива слов
		setWords = map[string]bool{}                               // инициализация мапы для записи слов
	)

	for _, item := range words { // добавление элементов в мапу, где ключ - слово из массива, значение может быть любым
		setWords[item] = true // тк мапа не может содержать одинаковые ключи, они перезаписываются, в результате получатся
		// уникальные ключи без повторений
	}

	for word := range setWords {
		fmt.Println(word) // вывод на консоль результата
	}
}
