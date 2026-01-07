package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumOfDigits(56))
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

/*

Detailed Dry Run:
------------------------

Let's step through exactly what happens when main() runs and calls sumOfDigits(56):

1. main() function is called.
2. Inside main(), fmt.Println(sumOfDigits(56)) is executed.
   - Before Println can print, sumOfDigits(56) must be computed.

Step-by-step trace of sumOfDigits(56):

A. sumOfDigits(56)
   - n = 56, which is >= 10
   - So, return (56 % 10) + sumOfDigits(56 / 10)
   - (56 % 10) = 6
   - (56 / 10) = 5
   - So, return 6 + sumOfDigits(5)

B. sumOfDigits(5)
   - n = 5, which is < 10 (base case)
   - So, return 5

Back to previous call:

A. sumOfDigits(56) resumes:
   - It was waiting for 6 + sumOfDigits(5)
   - sumOfDigits(5) just returned 5
   - 6 + 5 = 11

3. The final result: sumOfDigits(56) returns 11 to main()

4. fmt.Println(11) prints:
   11

Summary Table:
------------------------
Call               Returns
sumOfDigits(5)     5
sumOfDigits(56)    6 + 5 = 11

------------------------
This function recursively sums the digits of the number.
In this example, 5 + 6 = 11.

*/

/* Another Example: sumOfDigits(1234)
Unwinding the recursion for 1234 for practice:

sumOfDigits(1234)
-> 1234 % 10 = 4, sumOfDigits(123)
sumOfDigits(123)
-> 123 % 10 = 3, sumOfDigits(12)
sumOfDigits(12)
-> 12 % 10 = 2, sumOfDigits(1)
sumOfDigits(1)
-> n = 1, base case, return 1

Working back up:
sumOfDigits(12) = 2 + 1 = 3
sumOfDigits(123) = 3 + 3 = 6
sumOfDigits(1234) = 4 + 6 = 10

So, sumOfDigits(1234) returns 10.
*/

