package main

import "fmt"

func main() {
	width := 2.5
	lenght := 3
	floatLenght := float64(lenght)
	fmt.Println("Area is", width * floatLenght)
	fmt.Println("width > lenght?", width > floatLenght)
}