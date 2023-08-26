package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	min_number, err := findMin(1, 2, 3, -8, 0)
	// min_number, err := findMin()
	if err == nil {
	fmt.Printf("Min number from slice: %d", min_number)
	} else {
		log.Println(err)
	}
}

func findMin(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("No numbers provided")
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min, nil
}