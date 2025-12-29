// This Go program demonstrates how functions in Go can return multiple values.
// Multiple return values are a powerful feature in Go, allowing functions to return 
// more than one result directly, without the need for complex data structures.
// This is often used to return a result along with an error, or to return several related values together.

package main

import (
	"errors"
	"fmt"
)

func main() {
	// Example 1: Multiple return values from divide()
	// Here, divide() returns both the quotient and remainder when dividing two numbers.
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
	// Output: Quotient: 3, Remainder: 1

	// Example 2: Multiple return values including an error value from compare()
	// In Go, it's common to return an error as an additional value to signal problems.
	var num1 int
	var num2 int
	fmt.Println("Enter two numbers to compare.")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	// compare() returns a string message and an error.
	result, err := compare(num1, num2)
	if err != nil {
		// If an error occurred, handle it.
		fmt.Println("Error:", err)
	} else {
		// Otherwise, use the result.
		fmt.Println(result)
	}
}

// divide demonstrates returning multiple named results: quotient and remainder.
// This is a common use case in Go, simplifying the return of multiple values.
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b   // quotient calculation
	remainder = a % b  // remainder calculation
	// When using named return values, a plain 'return' will return the current values.
	return
}

// compare demonstrates returning a value and a possible error, another common Go idiom.
// If the numbers are equal, it returns an error; otherwise, it returns a string result and nil error.
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater. Both numbers are equal.")
	}
}
