package main

import (
	"fmt"
	"strconv"
)

const errorMsg = "Error parsing value: "

func main() {
	// Example 1: Parse a decimal integer value from string using Atoi
	numStr := "12345"
	// strconv.Atoi parses a string as a base-10 integer and returns int type
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, numStr)
	} else {
		fmt.Printf("String %q parsed with strconv.Atoi: %d (type: %T)\n", numStr, num, num)
		// Demonstrate the value can be used as integer: increment by 1
		fmt.Printf("Value after incrementing: %d\n", num+1)
	}

	// Example 2: Parse using ParseInt which gives you int64. 10 is the base, 32 is the bit size.
	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, numStr)
	} else {
		fmt.Printf("String %q parsed with ParseInt (base 10, 32 bits): %d (type: %T)\n", numStr, numistr, numistr)
	}

	// Example 3: Parse a float from a string
	floatstr := "3.14"
	// ParseFloat returns a float64 value from the string
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, floatstr)
	} else {
		fmt.Printf("String %q parsed as float64: %.2f\n", floatstr, floatval)
	}

	// Example 4: Parse a binary string to decimal integer
	binarStr := "1010"
	decimal, err := strconv.ParseInt(binarStr, 2, 64)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, binarStr)
		return
	}
	fmt.Printf("Binary string %q parsed as decimal: %d\n", binarStr, decimal)

	// Example 5: Parse a hexadecimal string to decimal integer
	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, hexStr)
		return
	}
	fmt.Printf("Hexadecimal string %q parsed as decimal: %d\n", hexStr, hex)

	// Example 6: Attempt to parse an invalid integer string
	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Printf("%s %v (input was: %q)\n", errorMsg, err, invalidNum)
		// Demonstrate that when parsing fails, you usually want to handle or report the error
		return
	}
	fmt.Printf("String %q parsed as integer: %d\n", invalidNum, invalidParse)
}
