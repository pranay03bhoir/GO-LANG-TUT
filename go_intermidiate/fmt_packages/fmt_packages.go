// Go Fmt Package Demo: Detailed Explanations and Improved Output

/*
	This program demonstrates the use of major printing,
	formatting, scanning, and error formatting functions from the Go "fmt" package.
	It not only executes the code but also provides clear, step-by-step commentary,
	and improves the output with informative print statements so you can see *why* and *how* each line works.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("==== Go 'fmt' Package Demonstration with Detailed Explanations ====\n")

	// ---- 1. Basic Printing Functions ----
	fmt.Printf(">>> Section 1: Simple Print, Println, and Printf <<<\n")

	fmt.Println("- fmt.Print prints its arguments without any spaces or newline:")
	fmt.Print("Hello ")
	fmt.Print("World!") // No newlines, result: "Hello World!"
	fmt.Print(12, 456)  // No separator between arguments, but inserts default spaces
	fmt.Print("\n")     // Explicitly make a newline for clarity

	fmt.Println("\n- fmt.Println prints arguments separated by spaces, and adds a newline at the end:")
	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	fmt.Println("- You can use fmt.Printf for formatted output (like C's printf):")
	name := "Pranay"
	age := 25
	fmt.Printf("My name is %s & my age is %d\n", name, age)
	fmt.Printf("Show formatting: Binary: %b | Hex: %x (for age value)\n\n", age, age)
	fmt.Printf(">>> --------------------------------------------------------\n")

	// ---- 2. Build String with Formatting Functions ----
	fmt.Printf(">>> Section 2: String Formatting (Sprint, Sprintln, Sprintf) <<<\n")

	fmt.Println("- fmt.Sprint concatenates all arguments (like Print), but RETURNS a string instead of printing it.")
	s := fmt.Sprint("Hello", "World", 123, 456)
	fmt.Printf("fmt.Sprint(\"Hello\", \"World\", 123, 456): %q\n\n", s) // %q prints string with quotes for clarity

	fmt.Println("- fmt.Sprintln is like Sprint, but inserts spaces between arguments and appends a newline at the end (in the string):")
	sln := fmt.Sprintln("Hello", "World", 123, 456)
	fmt.Printf("fmt.Sprintln(\"Hello\", \"World\", 123, 456): %q (note it has \\n)\n", sln)
	fmt.Print("Output if printed with fmt.Print(...):\n")
	fmt.Print(sln)
	fmt.Print(sln) // Demonstrates printing twice (shows double spacing between, due to \n)

	fmt.Println("- fmt.Sprintf lets you compose a string using format verbs, just like Printf, but returns the formatted string.")
	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println("Sprintf result:", sf)
	fmt.Println("Printing Sprintf result again:", sf)
	fmt.Printf(">>> --------------------------------------------------------\n")

	// ---- 3. Scanning/Input Functions ----
	fmt.Printf(">>> Section 3: Reading User Input (Scan, Scanln, Scanf) <<<\n")

	var name1 string
	var age1 int

	fmt.Print("Enter your name and age, separated by a space (e.g.: John 34): ")

	// NOTE: fmt.Scan, Scanln, and Scanf all read input into provided variables.
	// - Scan: reads space-separated values into successive arguments.
	// - Scanln: like Scan, but stops at newline.
	// - Scanf: reads formatted input (like Printf, but with pointers)
	// We'll use Scanf for more control:
	n, errScan := fmt.Scanf("%s %d", &name1, &age1)
	fmt.Printf("Values successfully scanned: %d, Error: %v\n", n, errScan)
	fmt.Printf("You entered - Name: %s, Age: %d\n", name1, age1)
	fmt.Printf(">>> --------------------------------------------------------\n")

	// ---- 4. Error Formatting ----
	fmt.Printf(">>> Section 4: Formatting Errors with fmt.Errorf <<<\n")
	fmt.Println("We'll now use checkAge(age1) to validate if user is old enough to drive.")

	err := checkAge(age1)
	if err != nil {
		fmt.Println("Error returned from checkAge:", err)
	} else {
		fmt.Println("Age check passed! You are old enough to drive.")
	}

	fmt.Println("\n==== End of Demonstration ====")
}

// checkAge returns an error if age1 is less than 18; otherwise nil.
func checkAge(age1 int) error {
	fmt.Printf("(checkAge called with age1 = %d)\n", age1)
	if age1 < 18 {
		// fmt.Errorf creates a formatted error message as an error object
		return fmt.Errorf("Age %d is too young to drive.", age1)
	}
	return nil
}
