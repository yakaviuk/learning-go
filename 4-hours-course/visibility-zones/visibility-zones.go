package main

import (
	"fmt"
)

var a, b, c int = 4, 5, 6

func main() {
	a, b, c := 1, 2, 3
	a, b, c = c, a, b
	fmt.Println(a, b, c)

	myprint()
}

func myprint() {
	// a, b, c := 50, 100, 15
	a, b, c = c, a, b
	fmt.Println(a, b, c)
}
