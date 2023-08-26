package main

import "fmt"

func main() {

	// messages := []string{
	// 	"message 1",
	// 	"message 2",
	// 	"message 3",
	// 	"message 4",
	// }
	// messages[4] = "message 5"

	// panic("SOS!")
	// fmt.Print("really?")

	printMessage()
	fmt.Println("main()")

}

func printMessage() {
	defer printAgain()
	fmt.Println("printMessage()")

}

func printAgain() {
	fmt.Println("printAgain()")
}
