package main

import "fmt"

func main() {
	var name string = "Mick"
	var age int = 15
	fmt.Println(name, age)
	message, _ := enterTheClub(age)
	fmt.Println(message)
	//	fmt.Println(entered)
	// fmt.Println(enterTheClub(age))

}

func sayMore(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "You are allowed to enter", true
	}
	return "You arent't allowed to enter, go home", false

}
