package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(now)
	fmt.Println(year)

	var timeString string = now.String()
	fmt.Println(timeString)

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

}