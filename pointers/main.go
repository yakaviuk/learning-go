package main

import "fmt"

func main() {
	message := "I will become Ninja soon!"
	printMsg(message)

	fmt.Println(message)
}

func printMsg(message string) {
	message += " (from printMsg() function)"
	fmt.Println(message)
}
