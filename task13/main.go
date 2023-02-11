package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		box = "box"
		cat = "cat"
	)

	log.Printf("box = [%s], cat = [%s]", box, cat)
	fmt.Printf("%s is front of the %s\n", cat, box)

	box, cat = cat, box // тут происходит обмен значениями

	log.Printf("box = [%s], cat = [%s]", box, cat)
	fmt.Printf("%s is front of the %s\n", cat, box)
}
