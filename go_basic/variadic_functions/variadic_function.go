// This Go program demonstrates the use of variadic functions.
// A variadic function is a function that can accept zero or more arguments of a specified type.
// Variadic parameters are often used when you want a function to handle a variable number of inputs.
//
// In Go, a variadic parameter is defined by putting '...' before the type of the last parameter in a function definition,
// for example: func sum(nums ...int).
//
// This allows you to call the function with any number of int arguments,
// and inside the function, 'nums' is treated as a slice of ints.

package main

import "fmt"

func main() {
	// The sum function is a variadic function, which means it can accept any number of integers after the first string argument.
	// The first argument here is a string statement that will be returned alongside the sum.
	statement, total := sum("The sum of 1,2,3 is:", 1, 2, 3)
	fmt.Println(statement, total) // Output: The sum of 1,2,3 is: 6

	// You can also pass a slice of integers to a variadic function using the ... syntax.
	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence, total := add(3, numbers...)
	fmt.Println("Sequence:", sequence, "total:", total) // Output: Sequence: 3 total: 24
}

// sum is a variadic function.
// It takes a string statement and any number of int arguments (nums ...int).
// The function sums up all integers provided and returns the statement and the total.
func sum(statement string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return statement, total
}

// add is another example of a variadic function.
// It takes an initial integer sequence and then any number of more int arguments.
// The function sums up all the ints in nums and returns both sequence and the total sum.
func add(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
