package main

import "fmt"

func main() {
	printMessage("foo")
	printMessage("bar")
	printMessage("baz")

	mickHello := sayHello("Mick")
	printMessage(mickHello)

	printMessage(sayMore("Mick", 42))
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string) string {
	return "Hello, " + name
}

func sayMore(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
	// return "Hello, " + name + ". You are " + age + " years old."
}
