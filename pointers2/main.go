package main

import "fmt"

func main() {
	message := "I will become Ninja soon"
	printMsg(&message) // referencing 

	fmt.Println(message)
}

func printMsg(message *string) { // *string - pointer
	*message += " (from printMsg() function)" // dereference
	fmt.Println(message)
}
