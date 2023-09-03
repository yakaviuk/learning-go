package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	double(&amount)
	fmt.Println("amount in main func:", amount)
	fmt.Println("amount vars address:", &amount)
	fmt.Println("type:", reflect.TypeOf(&amount))
}

func double(number *int) {
	*number *= 2
	fmt.Println("number in function:", *number)
	fmt.Println("number's address in function:", number)

}