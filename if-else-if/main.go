package main

import "fmt"

func main() {
	var name string = "Mick"
	var age int = 13
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
	if age >= 18 && age < 45 {
		return "You are allowed to enter", true
	} else if age >= 45 && age < 65 {
		return "Are you sure that you would like the music?", true
	} else if age >= 65 {
		return "It's too much for you", false
	}

	return "You arent't allowed to enter, go home", false

}
