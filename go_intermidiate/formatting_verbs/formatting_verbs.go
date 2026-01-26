package main

import (
	"fmt"
)

func main() {
	// ===== Detailed Go Formatting Verbs Demonstration =====
	// This code provides a thorough walkthrough of common formatting verbs used with fmt.Printf in Go.
	// For each section, not only do we demonstrate the formatting, but we also clearly print what is being shown.
	// This approach helps in understanding the output and connecting it directly to the formatting verb.

	// ---- Section 1: General Formatting Verbs ----
	// %v  : Prints the value in default format.
	// %#v : Prints the value in Go-syntax representation.
	// %T  : Prints the type of the value.
	// %%  : Prints a literal % sign.

	fmt.Println("=== General Formatting Verbs ===")
	num := 150_000.5 // Using underscores makes numbers more readable (Go 1.13+)
	text := "Hello World"

	fmt.Printf("Value of num (%%v): %v\n", num)
	fmt.Printf("Go-syntax value of num (%%#v): %#v\n", num)
	fmt.Printf("Type of num (%%T): %T\n", num)
	fmt.Printf("Value of num followed by percent sign (%%v%%%%): %v%%\n", num)

	fmt.Printf("Value of text (%%v): %v\n", text)
	fmt.Printf("Go-syntax value of text (%%#v): %#v\n", text)
	fmt.Printf("Type of text (%%T): %T\n", text)
	fmt.Printf("Value of text followed by percent sign (%%v%%%%): %v%%\n\n", text)

	// ---- Section 2: Integer Formatting Verbs ----
	// %b  : Binary (base 2)
	// %d  : Decimal (base 10)
	// %+d : Decimal with always showing a sign
	// %o  : Octal (base 8)
	// %O  : Octal with leading 0o (since Go 1.13)
	// %x  : Hexadecimal (base 16, lowercase)
	// %X  : Hexadecimal (base 16, uppercase)
	// %#x : Hexadecimal with leading 0x
	// %4d : Pad with spaces (width 4, right-justified)
	// %-4d: Pad with spaces (width 4, left-justified)
	// %04d: Pad with zeros (width 4)

	fmt.Println("=== Integer Formatting Verbs ===")
	numInt := 255
	fmt.Printf("numInt = %d\n", numInt)
	fmt.Printf("Binary (%%b): %b\n", numInt)
	fmt.Printf("Decimal (%%d): %d\n", numInt)
	fmt.Printf("Decimal, always show sign (%%+d): %+d\n", numInt)
	fmt.Printf("Octal (%%o): %o\n", numInt)
	fmt.Printf("Octal with leading 0o (%%O): %O\n", numInt)
	fmt.Printf("Hexadecimal lowercase (%%x): %x\n", numInt)
	fmt.Printf("Hexadecimal UPPERCASE (%%X): %X\n", numInt)
	fmt.Printf("Hexadecimal with leading 0x (%%#x): %#x\n", numInt)
	fmt.Printf("Width 4, right-aligned, pad with spaces (%%4d): '%4d'\n", numInt)
	fmt.Printf("Width 4, left-aligned, pad with spaces (%%-4d): '%-4d'\n", numInt)
	fmt.Printf("Width 4, pad with zeros (%%04d): '%04d'\n\n", numInt)

	// ---- Section 3: String Formatting Verbs ----
	// %s  : Prints as plain string
	// %q  : Prints as double-quoted string (escaped)
	// %8s : Plain string, width 8, right-justified
	// %-8s: Plain string, width 8, left-justified
	// %x  : Hexadecimal dump of byte values (lowercase)
	// % x : Same as %x, but with spaces between bytes

	fmt.Println("=== String Formatting Verbs ===")
	exampleStr := "Pranay"
	fmt.Printf("Plain string (%%s): %s\n", exampleStr)
	fmt.Printf("Quoted string (%%q): %q\n", exampleStr)
	fmt.Printf("Right-aligned, width 8 (%%8s): '%8s'\n", exampleStr)
	fmt.Printf("Left-aligned, width 8 (%%-8s): '%-8s'\n", exampleStr)
	fmt.Printf("Hex dump of string bytes (%%x): %x\n", exampleStr)
	fmt.Printf("Hex dump with spaces (%% x): % x\n\n", exampleStr)

	// ---- Section 4: Boolean Formatting Verbs ----
	// %t : Value of the boolean as 'true' or 'false'
	// %v : Also outputs boolean in default (true/false) format

	fmt.Println("=== Boolean Formatting Verbs ===")
	boolTrue := true
	boolFalse := false
	fmt.Printf("Boolean true (%%t): %t\n", boolTrue)
	fmt.Printf("Boolean false (%%t): %t\n", boolFalse)
	fmt.Printf("Boolean false using %%v: %v\n\n", boolFalse)

	// ---- Section 5: Float Formatting Verbs ----
	// %e   : Scientific notation with 'e'
	// %f   : Decimal point, no exponent
	// %.2f : 2 decimal places
	// %6.2f: Width 6, 2 decimal places
	// %g   : Exponent as needed, minimum necessary digits

	fmt.Println("=== Float Formatting Verbs ===")
	floatNum := 9180078.18
	fmt.Printf("Scientific notation (%%e): %e\n", floatNum)
	fmt.Printf("Decimal notation (%%f): %f\n", floatNum)
	fmt.Printf("Precision of 2 (%%.2f): %.2f\n", floatNum)
	fmt.Printf("Width 6, prec. 2 (%%6.2f): '%6.2f'\n", floatNum)
	fmt.Printf("Compact (%%g): %g\n", floatNum)

	// ===== End of Verb Demonstration =====
	fmt.Println("\nEach above print is annotated with its corresponding formatting verb and what it does for utmost clarity.")
}
