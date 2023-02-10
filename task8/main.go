package main

import (
	"fmt"
	"log"
	"strconv"
)

func clearBit(number int64, position int) int64 {
	number = number &^ (1 << position) // это 10^(position - 1) например 1 << 3 = 1000
	// логическое и не там, где у 2-го числа 0 берется значение из 1-го числа, если у 2-го числа 1, у результата будет 0
	return number
}

func setBit(number int64, position int) int64 {
	number = number | (1 << position)
	// логическое или возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
	// поэтому только ноль на нужной позиции заменится на 1
	return number
}

func main() {
	var (
		number   int64 // объявление переменных для числа, сдвига, бинарного числа и результата
		position int
		binNum   string
		res      int64
	)

	log.Print("enter number and position: ")

	if _, err := fmt.Scan(&number, &position); err != nil { // получение нужных данных из консоли от пользователя
		log.Fatal(err.Error())
	}

	binNum = strconv.FormatInt(number, 2) // для того чтобы понять, существует ли указанный индекс в числе
	log.Printf("input num in binary system = %b\n", number)

	if len(binNum) < position { // проверка есть ли введенная позиция в числе
		log.Fatal("this binary number has no bit at position ", position)
	}

	if binNum[position] == 1 {
		res = clearBit(number, position)
	} else {
		res = setBit(number, position)
	}
	log.Printf("result = %d, binary result = %b", res, res)

}
