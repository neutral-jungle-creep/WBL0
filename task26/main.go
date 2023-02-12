package main

import (
	"fmt"
	"log"
	"strings"
)

func checkForUniqueness(line string) bool {
	var symbolsSet = map[rune]struct{}{}

	for _, symbol := range []rune(strings.ToLower(line)) { // все символы в один регистр, чтобы программа не считала их разными рунами
		if _, exists := symbolsSet[symbol]; exists {
			return false
		} else {
			symbolsSet[symbol] = struct{}{}
		}
	}
	return true
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil { // получение от пользователя строки
		log.Fatal(err.Error())
	}

	fmt.Printf("[%s] --- %v ", line, checkForUniqueness(line))
}
