package main

import (
	"fmt"
	"log"
	"strings"
)

func reverseLine(words []string) string {
	var result []string

	for i := len(words) - 1; i >= 0; i-- { // начиная с конца слайса
		result = append(result, words[i]) // добавление к результату элементов слайса
	}

	return strings.Join(result, " ") // объединение в строку с разделителем пробел
}

func main() {
	var wordLine = make([]string, 3, 3) // инициализация слайса для записи туда слов "snow dog sun — sun dog snow"
	// TODO - как записать в слайс из строк заранее неизвестное кол-во слов я пока не знаю
	for i := range wordLine {
		// в каждый элемент по индексу записывается значение с консоли, введенное через пробел
		if _, err := fmt.Scan(&wordLine[i]); err != nil {
			log.Fatal(err.Error())
		}
	}

	reversedLine := reverseLine(wordLine)
	fmt.Println(reversedLine)
}
