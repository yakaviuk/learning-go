package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var name string = "Mick"
	var age int = 13
	fmt.Println(sayMore(name, age))
	message, err := enterTheClub(age)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(message)
	//	fmt.Println(entered)
	// fmt.Println(enterTheClub(age))

}

func sayMore(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "You are allowed to enter", nil
	} else if age >= 45 && age < 65 {
		return "Are you sure that you would like the music?", nil
	} else if age >= 65 {
		return "It's too much for you", errors.New("you're too old")
	}

	return "You arent't allowed to enter, go home", errors.New("you're too young")
	// TODO: errors handling
}
