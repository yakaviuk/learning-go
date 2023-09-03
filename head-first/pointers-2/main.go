package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.6
	return &myFloat
}

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
}
