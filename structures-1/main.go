package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	height uint32
	weight float32
}

// constructor
func NewUser(name string, age int, sex string, height uint32, weight float32) User {
	return User{
		name:   name,
		age:    age,
		sex:    sex,
		height: height,
		weight: weight,
	}
}

func main() {
	user1 := User{"Tom", 52, "male", 183, 80.1}
	user2 := User{"Erik", 36, "male", 173, 70.1}
	user3 := User{"Monika", 23, "female", 162, 53}
	user4 := User{"Caroline", 37, "female", 175, 67.9}
	// using constructor
	user5 := NewUser("Jack", 40, "male", 178, 80.5)

	fmt.Printf("%+v\n", user4)
	fmt.Printf("%+v\n", user5)
	fmt.Printf("%+v\n", user3)
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user1)

	fmt.Println(user1.name, user1.age)
	printUserInfo(user3)
}

func printUserInfo(user User) {
	fmt.Println("name:", user.name, "sex:", user.sex)
}