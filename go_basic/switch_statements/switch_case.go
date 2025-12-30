package main

import (
	"fmt"
)

func main() {

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("It's an unknown fruit")
	}

	fmt.Println("")

	//Multiple condition switch case

	var day string
	fmt.Println("Type a day of the week")
	fmt.Scanln(&day)

	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("It's a weekday")
	case "sunday", "saturday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a invalid day")
	}

	var number int
	fmt.Println("Enter a number")
	fmt.Scanln(&number)

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("The number is 2")
	default:
		fmt.Println("Not Two")
	}

	checkType(2)
	checkType(3.14)
	checkType("pranay")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("unknown type")
	}
}
