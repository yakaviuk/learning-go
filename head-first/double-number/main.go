package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	double(amount)
	fmt.Println("number in main func:", amount)
	fmt.Println("vars address:", &amount)
	fmt.Println("type:", reflect.TypeOf(&amount))
}

func double(number int) {
	number *= 2
	fmt.Println("number in function:", number)
}