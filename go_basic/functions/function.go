// This Go program demonstrates various aspects of functions, including 
// basic function definitions, anonymous functions, assigning functions to variables,
// passing functions as arguments, and returning functions from other functions (higher-order functions).

package main

import (
	"fmt"
)

func main() {
	// Calling a simple named function 'add' which takes two integers and returns their sum.
	sum := add(2, 47)
	fmt.Println(sum) // Output: 49

	// Anonymous function: This function does not have a name and is called immediately.
	func() {
		fmt.Println("Hello Anonymous function ")
	}()

	// Assigning an anonymous function to a variable.
	// The variable 'greet' holds a function with no parameters.
	greet := func() {
		fmt.Println("Hello anonymous function stored in greet variable")
	}
	greet() // Calling the function via the variable

	// Assigning a named function to a variable.
	// Since functions are first-class citizens in Go, they can be assigned to variables.
	operation := add
	result := operation(4, 5) // Using the variable as a function
	fmt.Println(result)        // Output: 9

	// Passing a function as an argument to another function.
	// Here we call 'applyOperation', passing in two numbers and the 'add' function.
	newResult := applyOperation(4, 9, add)
	fmt.Println("applyOperation function called 4 + 9 : ", newResult) // Output: 13

	// Returning a function from another function (Higher-order function).
	// 'createMultiplier' returns a function that multiplies its input by the factor provided.
	multiplyByTwo := createMultiplier(2)                 // multiplyByTwo is a function: func(int) int
	fmt.Println("Multiply by two 6 * 2 : ", multiplyByTwo(6)) // Output: 12
}

// Simple named function: Adds two integers.
func add(a, b int) int {
	return a + b
}

// Higher-order function: Takes two integers and a function (that operates on two ints and returns int) as arguments.
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Higher-order function: Returns a function that multiplies its argument by the provided factor.
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
