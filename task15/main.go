package main

import (
	"fmt"
	"math/rand"
)

// 1 переменная в поле видимости всего пакета - это плохая практика
// 2 со строкой нужно работать как с массивом рун, это поможет избежать неожиданного результата при работе с unicode
// 3 имя переменной, функции или класса должно сообщить, почему эта переменная существует, что она делает и как
//используется, то есть имена someFunc и justString выбраны неправильно
// 4 объявление лишней переменной, можно упростить код до --- justString := createHugeString(1 << 10)[:100]

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

func makeRandSymbolsString() string {
	return string(makeRandRuneCollection(1 << 10)[:100])
}

func makeRandRuneCollection(length int) []rune {
	runeCollection := make([]rune, length)

	for i := range runeCollection {
		runeCollection[i] = rune(1 + rand.Intn(3))
	}

	return runeCollection
}

func main() {
	randomString := makeRandSymbolsString()

	fmt.Println(randomString)
}
