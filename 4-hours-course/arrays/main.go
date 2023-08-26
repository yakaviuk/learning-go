package main

import (
	"fmt"
	"errors"
)

func main() {

	messages := [3]string{"1", "2", "3"}
	printMessages(messages)
	fmt.Println(messages)
}

func printMessages(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty slice")
	}

	messages[1] = "5"
	fmt.Println(messages)

	return nil
}
