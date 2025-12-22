package main

import (
	"fmt"
)

/*
Example of how to use for loops in Go.

First, a simple iteration over a range from 1 to 5 is shown.

Next, an example of how to iterate over a collection is shown.

This example uses the range keyword to iterate over a slice of integers.

The first value of the range is the index of the value in the slice,
while the second value is the value itself.

The fmt package is used to print the values to the console.
*/
func main() {

	// Simple iteration over a range.
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// Iteration over a collection.
	// 1. Initialize a slice of integers.
	// In Go, slices are zero-indexed, meaning the first element is at position 0.
	numbers := []int{1, 2, 3, 4, 5, 6}
	// 2. The 'for range' loop:
	// 'index' receives the current position (0, 1, 2...)
	// 'values' receives a copy of the actual data at that position (1, 2, 3...)
	// 'range' tells Go to iterate through the entire 'numbers' slice.
	for index, values := range numbers {
		// 3. Print the results:
		// %d is a placeholder for an integer (decimal).
		// \n moves the cursor to a new line for the next iteration.
		fmt.Printf("Index: %d, Values: %d\n", index, values)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd numbers: ", i)
		if i == 5 {
			break
		}
	}

	rows := 5

	for i := 0; i < rows; i++ {
		for j := 1; j < rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// New for loop which iterates over a range of number
	// Example -
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("We have a lift off")
}
