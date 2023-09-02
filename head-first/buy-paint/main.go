package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64

func main() {
	metersPerLiter = 10.0
	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f liters needed for the first wall\n", amount)
		total += amount
	}
	amount, err = paintNeeded(-5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f liters needed for the second wall\n", amount)
		total += amount
	}
	fmt.Printf("In total %.2f paint liters needed", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / metersPerLiter, nil
}
