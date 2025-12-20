package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3

	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	var num float64 = 22.0 / 7.0
	fmt.Println(num)

	// overflow with signed integer
	var maxInt int64 = 9223372036854775807 // max value the int64 can handle.
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 // max value the uint64 can handle.
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}

