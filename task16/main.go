package main

import "fmt"

func quickSort(numsCollection []int) []int {
	if len(numsCollection) <= 1 {
		return numsCollection
	}

	var (
		elem   = numsCollection[0]
		left   = make([]int, 0, 5)
		right  = make([]int, 0, 5)
		center = make([]int, 0, 5)
	)

	for _, num := range numsCollection {
		if num < elem {
			left = append(left, num)
		} else if num > elem {
			right = append(right, num)
		} else {
			center = append(center, num)
		}
	}
	return append(append(quickSort(left), center...), quickSort(right)...)

}

func main() {
	result := quickSort([]int{2, 50, 23, 34, 5, 2, 87, 23, 54, 7, 4, 3, 2, 343, 423, 423, 42, 5, 46, 5464})

	fmt.Println(result)
}
