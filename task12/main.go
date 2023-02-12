package main

import "fmt"

func main() {
	var (
		words    = [...]string{"cat", "cat", "dog", "cat", "tree"}
		setWords = map[string]struct{}{} // мапа для записи слов, значение может быть любым, лучше сделать пустую структуру,
		// потому что она занимает 0 байт в памяти
	)

	for _, item := range words {
		setWords[item] = struct{}{} // тк мапа не может содержать одинаковые ключи, они перезаписываются, в результате получатся
		// уникальные ключи без повторений
	}

	for word := range setWords {
		fmt.Println(word)
	}
}
