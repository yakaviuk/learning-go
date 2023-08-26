package main

import "fmt"

func main() {
	var users map[string]int
	// users := map[string]int
	// users = map[string]int{}
	users = make(map[string]int)
	users["Tom"] = 5

	fmt.Println(users)

	fmt.Println(len(users))
}