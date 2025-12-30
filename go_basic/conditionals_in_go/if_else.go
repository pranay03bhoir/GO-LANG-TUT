package main

import (
	"fmt"
)

func main() {

	age := 25

	if age > 18 {

		fmt.Println("You are an Adult")
	} else {
		fmt.Println("You are not an Adult")
	}

	// Examples of if else-if else chaining

	score := 66

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// Nested if else

	var userInput int
	fmt.Println("Enter a number of your choice")
	fmt.Scanln(&userInput)
	if userInput%2 == 0 {
		if userInput%3 == 0 {
			fmt.Println("The number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is not divisible by 2 and 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}
}
