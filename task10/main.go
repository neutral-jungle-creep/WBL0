package main

import (
	"fmt"
)

func addTValuesToGroups(tValues []float32) map[int][]float32 {
	var result = map[int][]float32{} // инициализация мапы для записи результата, ключ - группа

	for _, item := range tValues {
		group := int(item) / 10 * 10                // тк результат деления int - всегда целочисленное
		result[group] = append(result[group], item) // добавление значения в мапу по ключу группы в конец списка со значением
	}
	return result
}

func main() {
	var tValues = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // инициализация массива со значениями температур

	groups := addTValuesToGroups(tValues) // группировка значений в мапу

	for i, item := range groups {
		fmt.Printf("%d:%v ", i, item) // вывод на консоль результата
	}
}
