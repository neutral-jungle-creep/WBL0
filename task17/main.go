package main

import (
	"log"
	"sort"
)

func binarySearch(sortedNumsCollection []int, item int) (int, int) {
	var (
		low     = 0
		high    = len(sortedNumsCollection) - 1
		counter int
	)

	for low <= high {
		counter++
		center := (low + high) / 2
		num := sortedNumsCollection[center]

		log.Printf("нижняя граница = [%d] --- верхняя граница = [%d] --- середина = [%d]",
			sortedNumsCollection[low], sortedNumsCollection[high], num)

		if item > num {
			low = center + 1
		}
		if item < num {
			high = center - 1
		}
		if item == num {
			return item, counter
		}
	}
	return item, counter
}

func main() {
	searchField := []int{2, 5, 8, 12, 3, 23, 100, 11, 72, 1}
	sort.Ints(searchField) // с не отсортированным списком бинарный поиск не будет эффективен

	num, counter := binarySearch(searchField, 10)
	log.Printf("number = [%d]\n", num)
	log.Printf("counter = [%d]", counter)
}
