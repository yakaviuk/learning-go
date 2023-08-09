package main

import "fmt"

func main() {
	printMessage(sayMore("Mick", 42))
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayMore(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "You are allowed to enter", true
	} else {
		return "You arent't allowed to enter, go home", false
	}
}
