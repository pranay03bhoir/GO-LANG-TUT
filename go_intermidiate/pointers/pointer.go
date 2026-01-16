package main

import (
	"fmt"
)

// main demonstrates pointer basics in Go.
func main() {
	var a int = 10                        // Declare integer variable 'a', set to 10.
	var ptr *int                          // Declare a pointer to int.
	ptr = &a                              // Assign the address of 'a' to ptr.

	// Let's print detailed information about 'a' and the pointer:
	fmt.Println("Initial value of a:", a)
	fmt.Println("Address of a (&a):", &a)
	fmt.Println("Pointer ptr holds (should be address of a):", ptr)
	fmt.Println("Value at the address ptr points to (*ptr):", *ptr)

	// Call a function that modifies the value via the pointer
	handlePointer(ptr)

	// Print again to show the effect of pointer manipulation
	fmt.Println("Value of a after handlePointer(ptr):", a)
	fmt.Println("Value at the address ptr points to after update (*ptr):", *ptr)
}

// handlePointer takes a pointer to int and increments the value it points to.
func handlePointer(ptr *int) {
	// *ptr dereferences the pointer, accessing the actual integer value.
	// Increment the value by 1.
	*ptr++
}

/*
Detailed Explanation:

Pointers in Go are variables that store the memory address of another variable.
- 'a' is a regular int with value 10.
- 'ptr' is a pointer to int (type: *int).
- 'ptr = &a' means ptr now points to the memory address of 'a'.
- Printing 'a' gives the value (10).
- Printing '&a' (or 'ptr') gives the memory address (e.g., 0xc000014088).
- Printing '*ptr' dereferences the pointer, yielding the value stored at that address (initially 10).
- The function 'handlePointer' receives the pointer, and '*ptr++' increments the value at that address.
  So after this function, 'a' becomes 11.
- Printing 'a' and '*ptr' after shows both are now 11, since they refer to the same piece of data in memory.

This illustrates how pointers can be used to modify variables within functions and how Go lets you work with addresses directly.
*/