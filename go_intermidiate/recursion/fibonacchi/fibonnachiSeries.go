// Fibonacci Series using Recursion in Go
// -------------------------------------------------------
// This program prints the first 10 numbers in the Fibonacci sequence using recursion.
// Let's explain, in detail, how it works and provide a dry run.

package main

import (
	"fmt"
)

// The main function runs first when the program starts.
func main() {
	// Loop from 0 to 9 (10 numbers)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println() // For a clean newline at end
}

// fibonacci recursively calculates the n-th Fibonacci number.
func fibonacci(n int) int {
	if n < 2 {
		return n // Base cases: fibonacci(0)=0, fibonacci(1)=1
	}
	// Recursive case: sum of two previous numbers
	return fibonacci(n-1) + fibonacci(n-2)
}

/*
----------------- DETAILED DRY RUN ------------------------

Let's trace what happens when we run this program.

main() is called:
- It loops "i" from 0 to 9, printing fibonacci(i).

Let's trace a few outputs in detail.

----------------
First, i = 0:
	fmt.Printf("%d ", fibonacci(0))
	--> Calls fibonacci(0)
		fibonacci(0) < 2, so returns 0.
		Prints: 0

Second, i = 1:
	fmt.Printf("%d ", fibonacci(1))
	--> Calls fibonacci(1)
		fibonacci(1) < 2, so returns 1.
		Prints: 1

Third, i = 2:
	fmt.Printf("%d ", fibonacci(2))
	--> Calls fibonacci(2)
		2 >= 2, so recursively computes fibonacci(1) + fibonacci(0):
			fibonacci(1) = 1 (base case)
			fibonacci(0) = 0 (base case)
		Sum: 1 + 0 = 1
		Prints: 1

Fourth, i = 3:
	fmt.Printf("%d ", fibonacci(3))
	--> Calls fibonacci(3)
		3 >= 2, so computes fibonacci(2) + fibonacci(1)
			fibonacci(2): calls fibonacci(1) + fibonacci(0)
				fibonacci(1) = 1 (base)
				fibonacci(0) = 0 (base)
				So fibonacci(2) = 1
			fibonacci(1): 1 (base)
		So fibonacci(3) = 1 + 1 = 2
		Prints: 2

Fifth, i = 4:
	fmt.Printf("%d ", fibonacci(4))
	--> Calls fibonacci(4)
		4 >= 2, so computes fibonacci(3) + fibonacci(2)
			fibonacci(3): calls fibonacci(2) + fibonacci(1)
				fibonacci(2): calls fibonacci(1) + fibonacci(0)
					fibonacci(1) = 1; fibonacci(0) = 0
					fibonacci(2) = 1
				fibonacci(1) = 1
				fibonacci(3) = 1 + 1 = 2
			fibonacci(2): calls fibonacci(1) + fibonacci(0)
				fibonacci(1) = 1; fibonacci(0) = 0
				fibonacci(2) = 1
		So fibonacci(4) = 2 + 1 = 3
		Prints: 3

This continues up to i = 9, recursively building up the Fibonacci sequence.

----------------
Summary Table
-------------------------------------------
i   fibonacci(i)
0   0
1   1
2   1
3   2
4   3
5   5
6   8
7   13
8   21
9   34

Output printed:
0 1 1 2 3 5 8 13 21 34 

---------------------------
Key Points:
- The recursion stops when n < 2.
- Each fibonacci(n) is the sum of the two previous Fibonacci numbers.
- For each value of i, the recursive call stack is built down to the base cases (0 or 1), and then returns are combined up.
*/