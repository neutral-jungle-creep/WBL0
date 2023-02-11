package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		numOne = new(big.Int) // объявление переменных с помочью функции new потому что она возвращает ссылку на тип
		numTwo = new(big.Int)
		result = new(big.Int)
	)

	numOne.SetString("2147483648", 10) // 2^31 заполнение чисел значениями
	numTwo.SetString("1073741824", 10) // 2^30
	fmt.Printf("num 1 = [%d] --- num 2 = [%d]\n", numOne, numTwo)

	// вывод результата
	fmt.Printf("умножение = %d\n", result.Mul(numOne, numTwo))
	fmt.Printf("деление = %d\n", result.Div(numOne, numTwo))
	fmt.Printf("сложение = %d\n", result.Add(numOne, numTwo))
	fmt.Printf("вычитание = %d\n", result.Sub(numOne, numTwo))
}
