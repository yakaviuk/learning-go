package main

import (
	"fmt"
)

func main() {
	foo, bar, baz := 1, 2, 3
	foo, bar = bar, foo
	fmt.Println(foo, bar, baz)

	myprint()
}

func myprint() {
	a, b, c := 50, 100, 15
	a, b, c = c, a, b
	fmt.Println(a, b, c)
}
