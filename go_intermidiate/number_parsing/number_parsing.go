package main

import (
	"fmt"
	"strconv"
)

const errorMsg = "Error parsing value: "

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println(errorMsg, err)
	}

	fmt.Println("Parsed number: ", num)
	// Checking if the string is actually converted to integer.
	fmt.Println("Parsed number: ", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(errorMsg, err)
	}

	fmt.Println("Parsed Integer: ", numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println(errorMsg, err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval)

	binarStr := "1010"
	decimal, err := strconv.ParseInt(binarStr, 2, 64)
	if err != nil {
		fmt.Println(errorMsg, err)
		return
	}
	fmt.Println("Parsed binary to decimal: ", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(errorMsg, err)
		return
	}
	fmt.Println("Parsed Hex to decimal: ", hex)

	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println(errorMsg, err)
		return
	}
	fmt.Println("Parsed Hex to decimal: ", invalidParse)

}
