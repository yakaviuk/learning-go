package main

import "fmt"

func main() {
	var x int
	for x = 1; x <= 43; x+=3 {
		fmt.Println(x)
	}
	fmt.Println("post loop")
	fmt.Println(x)
}
