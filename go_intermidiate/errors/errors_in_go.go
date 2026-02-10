package main

import (
	"errors"
	"fmt"
	"math"
)

// This file demonstrates different ways of creating and handling errors in Go.
// It covers three main ideas:
//  1) Returning simple errors from functions (see sqrt and process).
//  2) Creating your own error type that implements the error interface (see myError and processError).
//  3) Wrapping errors with additional context using fmt.Errorf and the %w verb (see readData and readConfig).

// sqrt returns the square root of x.
// If x is negative, it returns a non-nil error instead of a valid result.
//
// The convention in Go is:
//   result, err := someFunction(...)
//   if err != nil {
//       // handle the error path
//   }
//   // otherwise, use result on the success path
func sqrt(x float64) (float64, error) {
	// Check for invalid input.
	if x < 0 {
		// errors.New constructs a simple error value with a message.
		// Here we explain *what* went wrong.
		err := errors.New("math error: square root of negative number")

		// Printing inside a helper function is usually NOT done in real applications,
		// but here we do it to show *where* the error originates.
		fmt.Printf("sqrt: about to return an error: %v\n", err)

		// We return 0 (dummy value) plus the error.
		return 0, err
	}

	// No error: compute and return the real square root and a nil error.
	result := math.Sqrt(x)
	return result, nil
}

// process represents some work that can fail when given bad input.
// If the byte slice is empty, it returns an error; otherwise it returns nil (meaning success).
func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("error: empty dataset")
	}
	return nil
}

func main() {
	// ===============================
	// Example 1: Simple error from sqrt
	// ===============================
	fmt.Println("=== Example 1: simple error from sqrt ===")
	fmt.Print("Enter a number to find its square root: ")

	var userInp float64
	fmt.Scanln(&userInp)

	// Call sqrt and follow the common Go pattern: result + error.
	result, err := sqrt(userInp)
	if err != nil {
		fmt.Println("sqrt returned an error.")
		fmt.Printf("Error value: %v\n", err)
		// In a real program we might log the error or return it upwards.
		// For this example, we simply stop here.
		return
	}

	fmt.Println("sqrt succeeded without error.")
	fmt.Printf("Square root of %v is %v\n", userInp, result)

	// ===============================
	// Example 2: Function that errors on invalid input
	// ===============================
	fmt.Println("\n=== Example 2: function returning error when input is invalid ===")

	// Here we deliberately pass an empty slice to trigger the error path.
	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("process returned an error because the data slice was empty.")
		fmt.Printf("Error value: %v\n", err)
	} else {
		fmt.Println("process succeeded with non-empty data.")
	}

	// ===============================
	// Example 3: Custom error type
	// ===============================
	fmt.Println("\n=== Example 3: custom error type implementing the error interface ===")

	err1 := processError()
	if err1 != nil {
		fmt.Println("processError returned a custom error value.")
		fmt.Printf("Error value: %v\n", err1)
		fmt.Printf("Dynamic type of error: %T\n", err1)
		// Notice that even though the dynamic type is *myError,
		// we can still use it anywhere an error is expected.
	}

	// ===============================
	// Example 4: Wrapped error with additional context
	// ===============================
	fmt.Println("\n=== Example 4: wrapped error with additional context using %w ===")

	err2 := readData()
	if err2 != nil {
		fmt.Println("readData failed and returned a wrapped error.")
		fmt.Printf("Wrapped error value: %v\n", err2)

		// errors.Unwrap lets us recover the original error that was wrapped.
		original := errors.Unwrap(err2)
		fmt.Printf("Original (inner) error: %v\n", original)
		return
	}

	fmt.Println("Data read successfully (no error).")
}

// myError is a custom error type.
// Any type that has an Error() string method automatically satisfies the built-in error interface.
type myError struct {
	message string
}

// Error implements the error interface for *myError.
// The fmt package will call this method when you print the error.
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

// processError shows how you can return your own custom error type from a function.
// The declared return type is error, but the concrete value is *myError.
func processError() error {
	return &myError{"custom error message"}
}

// readData shows how to add extra context to an error.
// It calls readConfig, and if that fails, wraps the returned error with more information.
func readData() error {
	err := readConfig()
	if err != nil {
		// fmt.Errorf with %w wraps the original error value so that callers can still
		// inspect it using errors.Unwrap, errors.Is, or errors.As.
		return fmt.Errorf("readData: failed to read config: %w", err)
	}
	return nil
}

// readConfig simulates a lower-level function that can fail (for example, reading a file).
// Here it always fails so we can see how the error propagates up through readData and main.
func readConfig() error {
	return errors.New("config error: could not load configuration")
}

