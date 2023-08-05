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

}
