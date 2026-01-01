// This Go program demonstrates the use and behavior of the os.Exit() function
// in the context of deferred statements and function termination.
//
// The os.Exit() function terminates the program immediately with the provided
// exit code, and does *not* run deferred functions. This is in contrast to
// a panic, which will run deferred functions before the program terminates.

package main

import (
	"fmt"
	"os"
)

func main() {
	// This deferred statement will NOT be executed if os.Exit is called,
	// because os.Exit terminates the program IMMEDIATELY.
	defer fmt.Println("Defer statement")

	fmt.Println("Starting the main function")

	// os.Exit(1) terminates the program and returns status code 1 to the OS.
	// No code after this statement will run, including deferred statements.
	os.Exit(1)

	// The code below is unreachable and will not be executed.
	fmt.Println("End of main function")
}

// Detailed Explanation:
// 1. The program defines a deferred statement. Normally, deferred code executes when
//    the surrounding function (main) returns. However, os.Exit(1) does not perform any
//    stack unwinding -- it halts the program instantly, so "Defer statement" is NOT printed.
//
// 2. "Starting the main function" is printed first.
//
// 3. os.Exit(1) is called, the program terminates with exit code 1, and the deferred
//    print statement, as well as all lines following os.Exit, are never executed.
//
// In summary, os.Exit() provides a way to exit a Go program immediately with a status code,
// skipping all deferred function calls and making proper resource cleanup impossible.
