package main

import "fmt"

func main() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	inc2 := increment()
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())
	fmt.Println(inc2())
}

func increment() func() int {
	count :=0
	return func() int {
		count++
		return count
	}
}