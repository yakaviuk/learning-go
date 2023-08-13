package main

import "fmt"

func main() {
	//var messages []string
	// messages := make([]string, 5)
	// messages := []string{"1", "2", "3", "4", "5"}
	intSlice := make([]int, 5, 15)
	fmt.Println(intSlice)
	for i := 0; i <len(intSlice); i++ {
		intSlice[i] =  i+1
	}

	fmt.Println(intSlice)

	for i := 0; i < 40; i++ {
		fmt.Printf("length: %d, capacity: %d, address: %p\n", len(intSlice), cap(intSlice), intSlice)
		intSlice = append(intSlice, i)
	}

	// slice with strings

	var messages []string
	messages = make([]string, 3000)
	messages = append(messages, "6")
	fmt.Println(len(messages))
	fmt.Println(cap(messages))




}
