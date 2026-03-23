package main

import (
	"fmt"
	"regexp"
)

/*
This program is a comprehensive demonstration of Go's regexp (regular expression) package.
It shows how to:
- Compile and use regular expressions to match email patterns.
- Use capturing groups to extract date components from a string.
- Perform search/replace on strings with regular expressions.
- Use regex flags (like case-insensitive matching).

Each section uses well-commented and informative print statements so you can see what is happening at each stage.
*/

func main() {
	// Demonstrate string literal and raw string literal usage in Go.
	fmt.Println("============================================")
	fmt.Println("Demo: Go String Literals")
	fmt.Println("--------------------------------------------")
	fmt.Println(`1. Interpreted string literal:  Hello, "I Am Great"`)
	fmt.Println(`2. Raw string literal:         Hello, "I am Great"`)
	fmt.Println("")

	// 1. Email Validation using Regular Expressions
	fmt.Println("============================================")
	fmt.Println("Demo 1: Email Pattern Matching")
	fmt.Println("--------------------------------------------")
	emailPattern := `[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	fmt.Printf("Regex pattern to match emails: %q\n", emailPattern)
	re := regexp.MustCompile(emailPattern)

	email1 := "user787@gmail.com"
	email2 := "invalid_email"
	fmt.Printf("Test email: %q --> Valid? %v\n", email1, re.MatchString(email1))
	fmt.Printf("Test email: %q --> Valid? %v\n", email2, re.MatchString(email2))
	fmt.Println("")

	// 2. Using Capturing Groups to Extract Date Components (YYYY-MM-DD)
	fmt.Println("============================================")
	fmt.Println("Demo 2: Capturing Groups for Dates")
	fmt.Println("--------------------------------------------")
	datePattern := `(\d{4})-(\d{2})-(\d{2})`
	fmt.Printf("Regex pattern to capture date (YYYY-MM-DD): %q\n", datePattern)
	re = regexp.MustCompile(datePattern)

	date := "2026-03-23"
	fmt.Printf("Input date string: %q\n", date)
	submatches := re.FindStringSubmatch(date)
	/*
		submatches[0] = entire matched substring
		submatches[1] = captured year (YYYY)
		submatches[2] = captured month (MM)
		submatches[3] = captured day (DD)
	*/
	if len(submatches) > 0 {
		fmt.Printf("Full match : %q\n", submatches[0])
		fmt.Printf("Year       : %q\n", submatches[1])
		fmt.Printf("Month      : %q\n", submatches[2])
		fmt.Printf("Day        : %q\n", submatches[3])
	} else {
		fmt.Println("No match found!")
	}
	fmt.Println("")

	// 3. Replacing all vowels in a string with '*'
	fmt.Println("============================================")
	fmt.Println("Demo 3: Replace Vowels in a String")
	fmt.Println("--------------------------------------------")
	str := "Hello World"
	vowelPattern := `[aeiou]`
	fmt.Printf("Original string             : %q\n", str)
	fmt.Printf("Regex pattern to match vowels: %q\n", vowelPattern)
	re = regexp.MustCompile(vowelPattern)
	replaced := re.ReplaceAllString(str, "*")
	fmt.Printf("After replacing vowels with '*': %q\n", replaced)
	fmt.Println("")

	// 4. Case-insensitive matching using regex flags
	fmt.Println("============================================")
	fmt.Println("Demo 4: Case-insensitive Regex Matching")
	fmt.Println("--------------------------------------------")
	caseInsensitivePattern := `(?i)go`
	fmt.Printf("Regex pattern (case-insensitive): %q\n", caseInsensitivePattern)
	re = regexp.MustCompile(caseInsensitivePattern)

	text := "Golang is going great"
	fmt.Printf("Test string : %q\n", text)
	match := re.MatchString(text)
	fmt.Printf("Does the regex match '%s'? %v\n", text, match)
	fmt.Println("")

	fmt.Println("============================================")
	fmt.Println("End of Regular Expressions Demonstration.")
	fmt.Println("============================================")
}
