package main

import "fmt"

func main() {
	foo := 5.3
	bar := 3.0
	fmt.Printf("About one-third: %0.2f\n", foo/bar)

	resultString := fmt.Sprintf("About one-third: %0.2f\n", foo/bar)
	fmt.Printf(resultString)

	goodName := "mobile phone"
	discount := 5
	fmt.Printf("You are going to buy %s with %d%% discount\n", goodName, discount)
	goodName1 := "wheels"
	fmt.Printf("Foo discount for %s is %v%%", goodName1, foo)

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("-----------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Printf("%.2f\n", 99.9299394394)


}
