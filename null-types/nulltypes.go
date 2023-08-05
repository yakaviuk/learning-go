package main

import (
	"fmt"
)

func main() {
	var foo string
	fmt.Println(foo)

	var bar rune
	// rune = int32 (number between -2147483648 and 2147483647)
	bar = 2147483647
	fmt.Println(bar)

	var baz float32
	baz = -15.52
	fmt.Println(baz)

	// boolean: null value is false
	var foobool bool
	foobool = true
	fmt.Println(foobool)

	// byte = uint8
	var a byte = 65
	fmt.Printf("%c\n", a)
}
