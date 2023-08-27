package main

import "fmt"

func main() {
	// var int string = "integer"
	// fmt.Println(int)

	var foo int = 5  // main.go:9:10: int (variable of type string) is not a type
	fmt.Println(foo) 
}