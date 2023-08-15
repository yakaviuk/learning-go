package main

import "fmt"

func main() {
	printInterface("qwe")
	printInterface(1231)
}

func printInterface(i interface{}) {
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}
	fmt.Println(len(str))
}
