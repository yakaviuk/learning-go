package main

import (
	"fmt"
	"reflect"
)

func main() {
	// message := ""
	var message string
	message = "Type your new text here"
	fmt.Println(message)

	const message1 string = "hello here"
	fmt.Println(message1)

	/*
		next: numbers
	*/

	var myNumber int = 12
	fmt.Printf("type of \"var myNumber\": ")
	fmt.Println(reflect.TypeOf(myNumber))

	////

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
	fmt.Println("#######")
	// fmt.Println("//////////////////////")

}
