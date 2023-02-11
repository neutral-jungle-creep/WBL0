package main

import (
	"fmt"
)

//func removeElemByIndex(symbols []string, index int) {
//	symbols = append(symbols[:index], symbols[index+1:]...)
//}
// TODO - отдельной функцией не получилось, пока что не могу понять, почему не перезаписывает слайс

func main() {
	var (
		symbols = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
		index   = 5
	)

	fmt.Printf("after remove = %v\n", symbols)
	symbols = append(symbols[:index], symbols[index+1:]...) // записываем в переменную все элементы кроме последнего
	fmt.Printf("before remove = %v", symbols)
}
