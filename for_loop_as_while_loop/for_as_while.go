package main

import (
	"fmt"
)

func main() {
	// Using for loop as a while Loop
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number: ", num)
		num++
	}

}
