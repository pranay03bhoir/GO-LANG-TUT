// Detailed Dry Run of the Code
// code :
package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(5))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*


Dry Run:
--------------------------
Let's follow exactly what happens when main() runs:

1. main() function is called.
2. Inside main(), fmt.Println(factorial(5)) is executed.
   - Before Println can print anything, factorial(5) must be computed.

3. factorial(5) is called:
   - n = 5, not zero, so else branch executes.
   - Return value = 5 * factorial(4)

4. factorial(4) is called:
   - n = 4, not zero.
   - Return value = 4 * factorial(3)

5. factorial(3) is called:
   - n = 3, not zero.
   - Return value = 3 * factorial(2)

6. factorial(2) is called:
   - n = 2, not zero.
   - Return value = 2 * factorial(1)

7. factorial(1) is called:
   - n = 1, not zero.
   - Return value = 1 * factorial(0)

8. factorial(0) is called:
   - n = 0, so base case reached.
   - Return value = 1

Now the recursive stack starts to resolve:

9. factorial(0) returns 1, so factorial(1) returns 1 * 1 = 1
10. factorial(1) returns 1, so factorial(2) returns 2 * 1 = 2
11. factorial(2) returns 2, so factorial(3) returns 3 * 2 = 6
12. factorial(3) returns 6, so factorial(4) returns 4 * 6 = 24
13. factorial(4) returns 24, so factorial(5) returns 5 * 24 = 120

Result:
- factorial(5) returns 120 to main()
- fmt.Println(120) prints:
  120

Summary Table:
----------------------------------
Call            Returns
factorial(0)    1
factorial(1)    1 * 1 = 1
factorial(2)    2 * 1 = 2
factorial(3)    3 * 2 = 6
factorial(4)    4 * 6 = 24
factorial(5)    5 * 24 = 120

Output: 120

--------------------------
This is a classic recursive factorial calculation.
*/
