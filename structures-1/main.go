package main

import "fmt"

type User struct {
	name   string
	age    Age
	sex    string
	height uint32
	weight float32
}

type Age int

func (a Age) isAdult() bool {
	return a>=18
}

// u User - value receiver
func (u User) printUserInfo() {
	// will be changed only in this function, original value will not be changed
	u.age = 42
	fmt.Println("name:", u.name, "sex:", u.sex, "age:", u.age)
}

// u *User - pointer receiver
func (u *User) changeAndPrintUserInfo() {
	// original value will be changed as well
	u.age = 42
	fmt.Println("name:", u.name, "sex:", u.sex, "age:", u.age)
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}

// constructor
func NewUser(name string, age int, sex string, height uint32, weight float32) User {
	return User{
		name:   name,
		age:    Age(age),
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
	// printUserInfo(user3)
	user3.printUserInfo()
	user4.printUserInfo()

	fmt.Println(user3.age)
	fmt.Println(user4.age)

	user1.changeAndPrintUserInfo()
	user2.changeAndPrintUserInfo()
	fmt.Println(user1.age)
	fmt.Println(user2.age)

	fmt.Println(user5.getName())
	user5.setName("Roland")
	fmt.Println(user5.getName())

	fmt.Println(user2.getName(), "is adult:", user2.age.isAdult())
}
