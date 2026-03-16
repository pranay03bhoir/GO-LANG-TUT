package main

import (
	"fmt"
)

// This program demonstrates various string formatting features in Go,
// including integer formatting with padding, string alignment, and the
// difference between interpreted and raw string literals.

func main() {
	// Integer formatting with zero padding
	num := 424
	// %05d formats the number as a 5-digit integer, padding with zeros on the left if necessary
	fmt.Printf("Integer with zero padding (width 5): [%05d]\n", num)
	// Output: [00424]

	// String alignment and width formatting
	message := "Hello"
	messageWithoutSpace := "Hellohjgvjhsdgfjshdvsjk"

	// %10s right-aligns the string in a field width of 10 characters
	fmt.Printf("Right-aligned string (width 10):   |%10s|\n", message)
	// Output: |     Hello|

	fmt.Printf("Right-aligned longer string (width 10): |%10s|\n", messageWithoutSpace)
	// Since the string is longer than 10, it's printed in full with no truncation or padding
	// Output: |Hellohjgvjhsdgfjshdvsjk|

	// %-10s left-aligns the string in a field width of 10 characters
	fmt.Printf("Left-aligned string (width 10):    |%-10s|\n", message)
	// Output: |Hello     |

	// Difference between interpreted and raw string literals
	var messageTwo string = "Hello \nWorld"   // Interpreted string literal: \n is treated as a newline
	var messageThree string = `Hello \nWorld` // Raw string literal: \n is treated as two characters '\' and 'n'

	fmt.Println("Interpreted string with \\n (results in new line):")
	fmt.Println(messageTwo)
	fmt.Println("Raw string with \\n (prints literal backslash-n):")
	fmt.Println(messageThree)

	// Example of a raw string for SQL queries (not used further in this program)
	sqlQuery := `SELECT * FROM users WHERE age > 18`
	_ = sqlQuery // Mark as used to avoid compiler error

	// Summary:
	// - %05d, %10s, and %-10s are used for formatting numbers and strings.
	// - Interpreted string literals handle escape sequences,
	//   while raw string literals do not.
	// - Useful for displaying well-formatted output and avoiding confusion with literals.
}
