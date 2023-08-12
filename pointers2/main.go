package main

import "fmt"

func main() {
	message := "I will become Ninja soon"
	fmt.Println(message)
	fmt.Println(&message) // 0x1400011e010
	changeMsg(&message) // referencing

	fmt.Println(message)
}

func changeMsg(message *string) { // *string - pointer
	*message += " (from printMsg() function)" // dereference
	fmt.Println(*message)
}
