package main

import (
	"errors"
	"fmt"
)

func main() {

	messages := []string{"1", "2", "3"}
	printMessages(messages)

	fmt.Println(messages)

}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty slice")
	}

	messages[1] = "5"
	fmt.Println(messages)

	return nil
}
