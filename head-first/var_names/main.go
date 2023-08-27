package main

import "fmt"

func main() {
	// var int string = "integer"
	// fmt.Println(int)

	var foo int = 5  // main.go:9:10: int (variable of type string) is not a type
	fmt.Println(foo) 
	
	// -------------

	goo := 9
	bar := 7

	baz, goo := 3, 4

	bar, foobar := 67, 89
	fmt.Println(goo)
	fmt.Println(bar)
	fmt.Println(baz)

	fmt.Println(foobar)
	//fmt.Println(foobaz)

	tor, dor := bar, foobar
	fmt.Println("tor value is:", tor)
	fmt.Println("dor value is:", dor)

	tor, dor = dor, tor
	fmt.Println("tor value is:", tor) // 89
	fmt.Println("dor value is:", dor) // 67

	fmt.Println("Initializing der")
	tor, der := dor, tor
	fmt.Println("tor value is:", tor) // 67
	fmt.Println("der value is:", der) // 89
}