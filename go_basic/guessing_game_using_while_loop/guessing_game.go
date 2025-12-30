package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	// Generate random number between 1 and 100
	target := rand.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the guessing game!!!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	// Take the user input
	var guess int
	for {
		fmt.Println("Enter Your Guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println(target, " is the correct number")
			break;
		}else if guess < target {
			fmt.Println("Too low!, try guessing higher number")
		}else {
			fmt.Println("Too high!, try guessing a lower number.")
		}
	}
}
