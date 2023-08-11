package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := prediction("Wednesday")
	fmt.Println(message)
	if err != nil {
		log.Println(err)
	}
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "Monday":
		return "Let's have a great beginning of the week", nil
	case "Tuesday":
		return "It will be beautiful day", nil
	case "Wednesday":
		return "Something in the middle", nil
	case "Thursday":
		return "Hard day!", nil
	case "Friday":
		return "It will be easy", nil
	case "Saturday":
		return "It will be easy", nil
	case "Sunday":
		return "It will be easy", nil
	default:
		return "Invalid week day", errors.New("invalid value")
	}
}
