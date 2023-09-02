package main

import "fmt"

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
}


func paintNeeded (width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}