package main

import "fmt"

func main() {
	foo := 1.0
	bar := 3.0
	fmt.Printf("About one-third: %0.2f\n", foo/bar)
	
	resultString := fmt.Sprintf("About one-third: %0.2f\n", foo/bar) 
	fmt.Printf(resultString)

	goodName := "mobile phone"
	discount := 5
	fmt.Printf("You are going to buy %s with %d%% discount\n", goodName, discount)
	goodName1 := "wheels"
	fmt.Printf("Foo discount for %s is %v%%", goodName1, foo)


}