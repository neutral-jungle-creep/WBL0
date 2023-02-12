package main

import (
	"fmt"
	"log"
	"math"
)

type Point struct { // тип данных выбран для получения максимально точного числа, не знаю, насколько это правильно
	// может быть было бы лучше преобразовать тип в функции подсчета, потому что math.pow, math.sqrt работают с float64
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // функция - конструктор с приставкой new должна возвращать ссылку на объект
	return &Point{
		x: x,
		y: y,
	}
}

func countDistBetweenTwoPoints(points [2]Point) float64 {
	return math.Sqrt(math.Pow(points[0].x-points[1].x, 2) + math.Pow(points[0].y-points[1].y, 2))

}

func main() {
	var (
		x, y   float64
		points [2]Point
	)

	for i := range points {
		fmt.Printf("enter %d point x and y: ", i+1)

		if _, err := fmt.Scan(&x, &y); err != nil { // получение координат точек с консоли
			log.Fatal(err.Error())
		}
		points[i] = *NewPoint(x, y)
	}

	dist := countDistBetweenTwoPoints(points)

	fmt.Printf("Dist between %f %f = %f", points[0], points[1], dist)

}
