package main

import "fmt"

func main() {
	user := struct{
		name string
		age int
		sex string
		height uint32
		weight float32
	}{"Tom", 52, "male", 183, 80.1}

	fmt.Printf("%+v\n", user)
	}

