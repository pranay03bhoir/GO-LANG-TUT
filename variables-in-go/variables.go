package main

import "fmt"

// This values is declares outside any block,
// therefore it is a global variable.
var middleName string = "Tanaji"

func main() {
	// The variables inside the main function can only be used inside the main function.
	// because these variable are in a blocked scope which is the main function.
	var age int = 20           // integer variable declaration
	var name string = "Pranay" // string variable declaration and initialization

	fmt.Println("Integer datatype variable ", age)
	fmt.Println("String datatype variable ", name)

	middleName := "Mayor" // Variables are mutable

	fmt.Println(middleName)

	// Go specific method to declare variable
	// By using this methods we do not specify the var keyword and the variable type
	// These thing are inferred by the go compiler
	count := 10
	lastName := "Bhoir"
	fmt.Println("Count ", count, " Lastname ", lastName)
	// Default values of the variables/Data-types
	// Numeric types : 0
	// Boolean types : false
	// String types : ""
	// Pointer, slices, maps, functions, and structs : nil
}
