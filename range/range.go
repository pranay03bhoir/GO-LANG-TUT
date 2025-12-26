package main

import (
	"fmt"
)

func main() {
	// The string that we want to iterate over
	message := "Hello World"

	// Using the 'range' keyword to iterate over a string in Go.
	// When you use 'range' with a string, it iterates over the string's Unicode code points (runes), not bytes.
	// For each iteration:
	//   - 'i' is the index (byte position in the string) where the rune starts.
	//   - 'v' is the value of the rune (Unicode code point) at that position.
	//
	// This is important because characters like accented letters or emoji can span multiple bytes,
	// but 'range' will give you each logical character (rune) and its starting position in bytes.

	for i, v := range message {
		// Print the index (byte offset) and the rune value (as an integer).
		fmt.Println(i, v)
		// Print a more readable message showing both the byte index and the character (as a rune).
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	/*
		EXPLANATION:

		- 'range' is a keyword in Go that can be used in a 'for' loop to iterate over elements in various data structures.
		  For strings, 'range' iterates over Unicode code points (runes), not just bytes.

		- Syntax: for index, value := range <collection> { ... }

		- In the context of a string:
		   * index: the starting byte position of the current rune in the string.
		   * value: the Unicode code point (rune) at that position.

		- Example Output (for "Hello World"):
		    0 72           (H)
		    1 101          (e)
		    2 108          (l)
		    3 108          (l)
		    4 111          (o)
		    5 32           (space)
		    6 87           (W)
		    7 111          (o)
		    8 114          (r)
		    9 108          (l)
		    10 100         (d)

		- If your string contains multibyte Unicode characters, the index values will increment by the number of bytes in each character, not always by 1.

		- This idiom ('for i, v := range <string>') is commonly used to process or analyze each character in a string, handling even non-ASCII characters correctly.
	*/
}
