package main

import "fmt"

func main() {
	number := 5
	var p *int

	p = &number
	fmt.Println(p)
	fmt.Println(&number)
	fmt.Println(number)

	*p = 10

	fmt.Println(number)

	number = 7
	fmt.Println(number)

	*p = 15
	fmt.Println(number)

}
