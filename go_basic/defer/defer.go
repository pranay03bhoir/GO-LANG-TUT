// This program demonstrates the use of the 'defer' statement in Go.
//
// The 'defer' keyword in Go schedules a function call (or method call)
// to be run after the function completes, just before the function returns.
// Deferred calls are executed in last-in, first-out (LIFO) order.
//
// In this example, multiple 'defer' statements are used inside the 'process' function.
// Their interactions show both the order of execution and the capture behavior of parameters at the point of defer.

package main

import (
	"fmt"
)

func main() {
	process(10) // Calls the process function with argument 10
}

func process(i int) {
	// The value of 'i' at the time of defer is captured. At this point, i is 10.
	defer fmt.Println("Deffered i value: ", i)

	// These statements are also deferred, and will be executed in LIFO order when process() returns.
	defer fmt.Println("First Defer statement executed")
	defer fmt.Println("Second Defer statement executed")
	defer fmt.Println("Third Defer statement executed")

	// This increments 'i', but does not affect the value captured by the deferred print above.
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i: ", i)

	// When process() is about to return, all deferred calls execute in LIFO order:
	// Output just before return will be:
	// Third Defer statement executed
	// Second Defer statement executed
	// First Defer statement executed
	// Deffered i value:  10
	//
	// Note: The final value printed for 'Deffered i value' is 10, because
	// that was the value of 'i' when defer was called, NOT when the function ended.
}
