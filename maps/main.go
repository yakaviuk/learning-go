package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasili":     15,
		"Petr":       23,
		"Kostyantyn": 48,
	}

	age, exists := users["Kostyantyn"]
	if exists {
		fmt.Println("Kostyantyn", age)
	} else {
		fmt.Println("no Kostyantyn in the map")
	}

	age, exists = users["Roma"]
	fmt.Println(age, exists)

	// fmt.Println(users)

	delete(users, "Vasili") // delete element from the map
	users["Roma"] = 54
	for key, value := range users {
		fmt.Println(key, value)
	}

}
