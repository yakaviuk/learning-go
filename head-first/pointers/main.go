package main

import "fmt"

func main() {
	var myInt int
	myInt = 4
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println("pointer itself:", myIntPointer)
	fmt.Println("pointer value", *myIntPointer)
	*myIntPointer = 9
	fmt.Println("new poiner value", *myIntPointer)
	

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}
