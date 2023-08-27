package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds) // deprecated
	min, max := 1, 101
	randNum := rand.Intn(max-min) + min
	fmt.Println("Guess the number!")
	for i := 10; i > 0; i-- {
		fmt.Println("You've got", i, "tries left")
		fmt.Print("Try to guess a number between 1 and 100: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		// guess, err := strconv.ParseInt(input, 10, 0)
		// guessInt := int(guess)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < randNum {
			fmt.Println("Oops, your guess was LOW")
		} else if guess > randNum {
			fmt.Println("Oops, your guess was HIGH")
		} else {
			fmt.Println("Congrats! You win")
			break
		}

		if i == 1 {
			fmt.Println("Sorry. You didn't guess my number. It was:", randNum)
		}
		
	}
}
