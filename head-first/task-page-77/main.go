package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	fileInfo, err1 := os.Stat("hello-world.go")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(fileInfo.Size())
}
