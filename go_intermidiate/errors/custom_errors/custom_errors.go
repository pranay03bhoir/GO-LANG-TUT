/*
This program demonstrates the use of custom error types in Go, error wrapping for retaining context, and improved error printing for clarity.

- We define a `customError` struct that implements the error interface, allowing us to add error codes and custom messages along with the underlying (wrapped) error.
- The `doSomething` function calls `doSomethingElse` and, if that function returns an error, wraps it in a `customError` to provide more context.
- The `main` function demonstrates error handling and prints out detailed error information for understanding how custom error types propagate.
*/

package main

import (
	"errors"
	"fmt"
)

// customError is our custom error type, which holds an error code, a message, and a wrapped error.
type customError struct {
	code    int    // An integer representing the error code
	message string // A custom human-readable message
	err     error  // The underlying error we are wrapping
}

// Error implements the error interface for customError.
// It returns a formatted string showing all details of the error.
func (e *customError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("CustomError [code=%d]: %s | cause: %v", e.code, e.message, e.err)
	}
	return fmt.Sprintf("CustomError [code=%d]: %s", e.code, e.message)
}

// main is the entry point where we call our function that might return a custom error.
// Detailed printing is provided so you can observe what happens at each step.
func main() {
	fmt.Println("=== Demonstration of Custom Error Handling ===")

	// Call doSomething which uses custom errors.
	err := doSomething()
	if err != nil {
		fmt.Println("An error occurred in doSomething:")
		fmt.Printf("Full error (with context):\n  %v\n", err)

		// Using errors.As to extract the custom error for more detailed info.
		var ce *customError
		if errors.As(err, &ce) {
			fmt.Println("This is a customError. Here are its fields:")
			fmt.Printf("  code:    %d\n", ce.code)
			fmt.Printf("  message: %s\n", ce.message)
			fmt.Printf("  cause:   %v\n", ce.err)
		} else {
			fmt.Println("Error is NOT of customError type. Original error:")
			fmt.Printf("  %v\n", err)
		}

		// Show that code after return is not executed.
		return
		fmt.Println("This line will never be printed because of the return above.")
	}

	fmt.Println("Operation completed successfully (no errors returned).")
}

// doSomething calls doSomethingElse and, if an error occurs, wraps it with a customError.
func doSomething() error {
	fmt.Println("-> In doSomething: calling doSomethingElse()")
	err := doSomethingElse()
	if err != nil {
		fmt.Println("-> doSomethingElse returned an error, wrapping in customError...")
		return &customError{
			code:    500,
			message: "Something went wrong in doSomething",
			err:     err,
		}
	}
	fmt.Println("-> doSomethingElse succeeded (no error).")
	return nil
}

// doSomethingElse simply simulates a lower-level function that always fails.
func doSomethingElse() error {
	fmt.Println("-> In doSomethingElse: simulating an internal error.")
	return errors.New("internal error: something failed at a low level")
}
