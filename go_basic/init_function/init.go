// This Go program demonstrates the use of the `init` function in the context of package initialization.
//
// Detailed Explanation:
//
// - The `init` function in Go is a special function that is automatically executed when a package is initialized,
//   that is, before the execution of the main function or before another package imports this package.
// - A package can have multiple `init` functions (within the same file or across multiple files in the same package).
//   These functions will be called in the order in which they appear in the source code files.
//
// - In this code, there are three `init` functions:
//       1. The first `init` function prints "Initialization package1...."
//       2. The second `init` function prints "Initialization package2...."
//       3. The third `init` function prints "Initialization package3...."
//
// - When the program is executed, all three `init` functions will be invoked automatically (in the order shown), 
//   even before the `main` function runs. 
// - After all the `init` functions complete, the `main` function executes and prints "Inside main function".
//
// Output when running this program:
// Initialization package1....
// Initialization package2....
// Initialization package3....
// Inside main function

package main

import (
	"fmt"
)

// First init function: called before main()
func init() {
	fmt.Println("Initialization package1....")
}

// Second init function: called after the previous init
func init() {
	fmt.Println("Initialization package2....")
}

// Third init function: called after the previous inits
func init() {
	fmt.Println("Initialization package3....")
}

// main is executed after all init functions have run
func main() {
	fmt.Println("Inside main function")
}
