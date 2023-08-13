package main

import "fmt"

func main() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	// for i := range messages {
	// 	fmt.Println(messages[i])
	// }

	for _, value := range messages {
		fmt.Println(value)
	}
	
	// counter := 0
	// for {
	// 	if counter == 100 {
	// 		break
	// 	}

	// 	counter++
	// 	println(counter)
	// }


}
