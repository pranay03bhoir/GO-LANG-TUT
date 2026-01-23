package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// SECTION 1: String Literals and Escape Sequences

	// Double quotes allow escape sequences like \n (newline), \t (tab), etc.
	message := "Hello, \nGo" // \n creates a new line between Hello, and Go
	fmt.Println("Double-quoted string with escape sequence (\\n):")
	fmt.Println(message)
	fmt.Printf("Above, \\n creates a new line.\n")

	// Backtick (`) defines raw string literals. Escape sequences are not processed.
	rawMessage := `Hello, \n Go` // \n appears as plain text
	fmt.Println("Backtick-quoted raw string (\\n not interpreted):")
	fmt.Println(rawMessage)
	fmt.Printf("Above, \\n appears as a regular backslash and n.\n")

	// Other escape sequences
	message1 := "Hello,\tGo" // \t inserts a horizontal tab
	message2 := "Hello, \rGo" // \r is carriage return. For demonstration, we display its bytes too.

	fmt.Println("Double-quoted with tab (\\t), output inserts a tab character:")
	fmt.Println(message1)
	fmt.Printf("Bytes of message1: %v\n", []byte(message1))

	fmt.Println("\nDouble-quoted with carriage return (\\r):")
	fmt.Println(message2)
	fmt.Printf("Bytes of message2: %v\n", []byte(message2))

	fmt.Printf("\n=========================================\n")

	// SECTION 2: String Properties and Indexing

	// Strings in Go are UTF-8 encoded, and you can get their byte length using len()
	fmt.Printf("Variable message: %q\n", message)
	fmt.Printf("Length of message in bytes: %d\n", len(message))

	// To access a specific byte:
	fmt.Printf("First byte in message: %q, value: %d\n", message[0], message[0])
	fmt.Println("Note: message[0] gives the byte value, which is the ASCII/UTF-8 of 'H'.")

	fmt.Printf("\n=========================================\n")

	// SECTION 3: String Concatenation

	greeting := "Hello, "
	name := "pranay"
	msg := greeting + name
	fmt.Printf("Concatenated string (greeting + name): %q + %q = %q\n", greeting, name, msg)

	fmt.Printf("\n=========================================\n")

	// SECTION 4: Lexicographical String Comparison

	str1 := "Apple"  // ASCII: A=65
	str2 := "banana" // b=98
	str3 := "app"    // a=97
	str4 := "apple"  // a=97

	fmt.Printf("Is %q < %q? %v\n", str1, str2, str1 < str2)
	fmt.Printf("Is %q < %q? %v\n", str3, str1, str3 < str1)
	fmt.Printf("Is %q > %q? %v\n", str4, str1, str4 > str1)
	fmt.Printf("Is %q > %q? %v\n", str4, str3, str4 > str3)
	// Explanation for one comparison
	fmt.Println("Explanation: ASCII('A')=65, ASCII('b')=98; so \"Apple\" < \"banana\" is true.")

	fmt.Printf("\n=========================================\n")

	// SECTION 5: Iterating Through a String (Using for ... range)

	fmt.Printf("Iterating over each rune (character) in %q:\n", message)
	for i, char := range message {
		fmt.Printf("Character at byte index %d: '%c' (Unicode: %U, Decimal: %v)\n", i, char, char, char)
	}
	fmt.Println("Note: Index represents the byte position (not the character position for multi-byte characters).")

	fmt.Printf("\n=========================================\n")

	// SECTION 6: Rune Count vs Byte Length

	fmt.Printf("Byte length of greeting (%q): %d\n", greeting, len(greeting))
	fmt.Printf("Rune (character) count in greeting: %d\n", utf8.RuneCountInString(greeting))

	fmt.Printf("\n=========================================\n")

	// SECTION 7: String Immutability

	greetingWithName := greeting + name
	fmt.Printf("Original greeting: %q\n", greeting)
	fmt.Printf("Concatenated: %q\n", greetingWithName)
	greeting = "Pranay" // Reassignment creates a new string
	fmt.Printf("Greeting after reassignment: %q\n", greeting)
	fmt.Println("Strings are immutable: existing values do not change underlying strings but point to new ones.")

	fmt.Printf("\n=========================================\n")

	// SECTION 8: Runes (Unicode Codepoints)

	var ch rune = 'a'  // Rune literal, Unicode for lowercase 'a' (97)
	var jch rune = '日' // Rune literal, Kanji for 'sun'/'day'

	fmt.Printf("Rune variable ch: '%c', Unicode codepoint: %U (Decimal: %d)\n", ch, ch, ch)
	fmt.Printf("Rune variable jch: '%c', Unicode codepoint: %U (Decimal: %d)\n", jch, jch, jch)

	fmt.Printf("\nPrinting runes as characters:\n")
	fmt.Printf("ch as character: %c\n", ch)
	fmt.Printf("jch as character: %c\n", jch)

	fmt.Println()

	// Converting a rune to a string results in a single-character string:
	cstr := string(ch)
	fmt.Printf("Converting rune ch (%c) to string: %q\n", ch, cstr)
	fmt.Printf("Type of cstr: %T\n", cstr)

	fmt.Printf("\n=========================================\n")

	// SECTION 9: Unicode Strings and Runes

	const NIHON = "こんにちは世界" // "Hello World" in Japanese

	fmt.Printf("Constant NIHON: %q\n", NIHON)
	fmt.Printf("Printing each character and its Unicode codepoint in \"こんにちは\":\n")
	jHello := "こんにちは"
	for i, runeValue := range jHello {
		fmt.Printf("Character %d: '%c' (Unicode: %U, Decimal: %d)\n", i, runeValue, runeValue, runeValue)
	}

	fmt.Printf("\n=========================================\n")

	// SECTION 10: Emoticons / Emojis in Go Strings

	emojis := "😂😂"
	fmt.Printf("Emoji string: %q\n", emojis)
	fmt.Printf("Bytes of emojis: %v\n", []byte(emojis))
	fmt.Printf("Printing emojis as %v: %v\n", "string (%%v)", emojis)
	fmt.Printf("Printing emojis as runes: ")
	for i, r := range emojis {
		fmt.Printf("[%d: %c (U+%04X)] ", i, r, r)
	}
	fmt.Println()
}
