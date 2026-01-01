// This Go program demonstrates the usage of the `recover` function for handling panics gracefully.
// In Go, a panic typically causes the program to terminate abruptly, but by using `recover` inside a deferred function,
// it is possible to intercept the panic and allow the program to continue executing.

package main

import "fmt"

func main() {
	// Call the process function, which might panic,
	// but any panic will be recovered internally.
	process()
	fmt.Println("Returned from process")
}

// process demonstrates how to use defer and recover to handle a panic condition gracefully.
// If a panic occurs in this function, the deferred anonymous function runs and recovers from the panic,
// preventing the program from terminating.
func process() {
	// Setup a deferred function that will run when the surrounding function returns
	// (either normally or because of a panic).
	defer func() {
		// recover() stops the panic and returns the error value passed to panic,
		// or nil if there is no ongoing panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("Start process")
	// Force a panic, simulating a run-time error.
	panic("Something went wrong")
	// The following line will not execute, as the control will jump to the deferred function
	// when panic occurs.
	fmt.Println("End process")
}

// Output of running this program:
// Start process
// Recovered: Something went wrong
// Returned from process
//
// Explanation:
// - The `process` function prints "Start process", then panics with the message "Something went wrong".
// - When panic occurs, all deferred functions are executed (here, the defer recovers from the panic).
// - Because recover is called, the panic is intercepted and the program continues normally after `process()`.
// - "Returned from process" is printed in main, showing the application did not crash.
