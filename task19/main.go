package main

import (
	"fmt"
	"log"
)

func reverseString(str string) string {
	var (
		runeLine = []rune(str) // создание списка рун из строки, так как посимвольно строку перебирать нельзя
		// символы unicode занимают 2 байта и длина строки может не равняться кол-ву символов в ней
		reverserStr string // объявление переменной для записи результата
	)

	for i := len(runeLine) - 1; i >= 0; i-- { // счетчик устанавливается на индекс последнего элемента списка рун
		reverserStr += string(runeLine[i]) // к результату прибавляется элемент списка рун от индекса, преобразованный в строку
	}

	return reverserStr
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil { // получение от пользователя строки
		log.Fatal(err.Error())
	}

	rLine := reverseString(line)

	fmt.Printf("reversed line = %s", rLine)
}
