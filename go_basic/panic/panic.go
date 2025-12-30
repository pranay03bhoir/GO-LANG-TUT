// This Go program demonstrates the use of panic and defer in error handling.
// The `panic` function is used to cause a run-time error that stops the ordinary flow of control.
// Any deferred functions will be executed before the program crashes due to panic.
// This example helps visualize how deferred functions work in the presence of a panic.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number of your choice ")
	var num int
	fmt.Scanln(&num)
	process(num)
}

// process demonstrates how deferred statements are executed even if a panic occurs.
func process(input int) {
	// Deferred functions are pushed onto a stack and are executed in LIFO order when the function returns,
	// either normally or via a panic.
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		// Trigger a panic if the input is negative.
		// At this point, all remaining deferred calls will run before the program terminates.
		panic("Input must be a non-negative number")
		// The following line will never be executed after panic.
		// fmt.Println("After panic")
	}
	fmt.Println("Processing input: ", input)
}
