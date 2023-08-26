package main

import "fmt"

func main() {
	var name string
	// var length, width float64 = 4.5, 9.8
	var length, width = 4.5, 9.8
	var quantity int

	name = "Erik Velaskes"
	// length, width 
	quantity = 1

	fmt.Println(name, "has ordered", quantity, "place.\nArea of the place is", length*width, "square meters")

	// var quantity float64 = 12.0 // head-first/vars/main.go:16:6: quantity redeclared in this block
	// fmt.Println(quantity)	// head-first/vars/main.go:8:6: other declaration of quantity

	fmt.Println("---------------------")
	
	// short form:
	another_name := "Lorenso Dalvatone"
	length1, width1 := 5.1, 8.01
	quantity1 := 4

	fmt.Println(another_name, "has ordered", quantity1, "places.\nArea of the each place is", length1*width1, "square meters")
}
