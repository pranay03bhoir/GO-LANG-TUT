package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"

	hello := "Hello"
	world := "World"

	fmt.Println(len(str))
	fmt.Println(hello + " " + world)
	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// Standard library functions :

	// String conversions

	var num int = 18
	var numToString string = strconv.Itoa(num)
	fmt.Println(len(numToString))
	// fmt.Println(len(num))

	// String splitting

	fruits := "apple,orange,banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)
	fruits1 := "apple-orange-banana"
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts1)

	// String joining

	countries := []string{"Germany", "France", "Italy", "India"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	// String contains

	fmt.Println(strings.Contains(str, "Go")) // It checks if the string contains the given values of string, //In our case "Hello Go!" contains "Go" ?. //The function returns true.

	// String replace and trim whitespace

	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone!"
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// String uppercase and lowercase

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	// String repeat and count

	fmt.Println(strings.Repeat("foo ", 3))
	fmt.Println(strings.Count("Hello", "l"))

	// String hasPrefix

	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasPrefix("Hello", "he"))

	// String hasSuffix

	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "la"))

	// Advanced string functions :

	// Regular expressions

	str1 := "He1llo, 123 Go 111!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str1, -1)
	fmt.Println(matches)

	str2 := "Hello ふしだらな女"
	fmt.Println(utf8.RuneCountInString(str2))
	// ==============================================================

	// STRING BUILDER

	var builder strings.Builder

	// write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// convert the builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using WriteRune() to add a character

	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder

	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
