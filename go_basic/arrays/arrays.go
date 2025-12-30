package main

import (
	"fmt"
)

func main() {

	var array [5]int // This is a blank initialization of an array.
	// so by default the value of this array will ber [0,0,0,0,0]

	fmt.Println(array)

	array[4] = 30 // This will insert the element 30 in the array at the 4th position.
	fmt.Println(array)

	array[0] = 22 //  This will insert the element 22 in the array at the 1st position.
	fmt.Println(array)

	fruits := [4]string{"Apple", "Banana", "Carrot", "Orange"}
	fmt.Println("The fruits array :", fruits)

	// In Go, assigning an array to another variable creates a copy of the array. Modifying the new variable does not affect the original array.

	// Demonstration of array assignment and how it creates a copy in Go

	originalArray := [3]int{1, 2, 3} // Create an array with 3 integers
	copiedArray := originalArray     // Assign originalArray to copiedArray; this creates a copy (not a reference)

	copiedArray[0] = 100 // Modify the first element of the copied array

	// Print both arrays to show that only copiedArray changed; originalArray remains unchanged
	fmt.Println("Original Array: ", originalArray) // Output: [1 2 3]
	fmt.Println("Copied Array: ", copiedArray)     // Output: [100 2 3]

	// Iterating over an array using for-loop

	for i := 0; i < len(array); i++ {
		fmt.Println("Element at index ,", i, ":", array[i])
	}

	// Iterating over an array using the range keyword in for loop

	for index, value := range array {
		fmt.Printf("Index: %d, Value:%d\n", index, value)
	}

	// Here, someFunction() returns two integer values: 1 and 2.
	// We use the short variable declaration with an underscore (_) to ignore the second value.
	a, _ := someFunction() // 'a' gets the first value (1), second value is discarded.
	fmt.Println(a)         // Output will be: 1

	// If you wanted to use both return values, you could write:
	// a, b := someFunction()
	// fmt.Println(a, b) // Would print: 1 2

	// Finding the length of the arrays using the len() function
	// Example -
	// The len() function in Go returns the number of elements in an array.
	// Here, we use len(array) to get the size of the array 'array', which is 5,
	// and print it to the console.
	fmt.Println("The length of the array is: ", len(array))

	// Comparing arrays
	// Example -
	// Here we declare two arrays, array1 and array2, both of type [3]int and containing the same integer elements: 1, 2, and 3.
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	// In Go, arrays can be compared using the '==' operator, but only if they are of the same type (length and element type must match).
	// This comparison checks whether all corresponding elements in both arrays are equal.
	// Since array1 and array2 have identical elements, this will print 'true'.
	fmt.Println("Array 1 is equal to Array 2 :", array1 == array2)

	// Multidimensional arrays
	// Example -
	// Creating a multidimensional array (a 3x3 matrix) in Go.
	// This matrix is an array of 3 arrays, each containing 3 integers.
	// The elements are organized in rows and columns as follows:
	// [ [1, 2, 3],
	//   [4, 5, 6],
	//   [7, 8, 9] ]
	matrix := [3][3]int{
		{1, 2, 3}, // first row
		{4, 5, 6}, // second row
		{7, 8, 9}, // third row
	}

	// Printing the entire matrix.
	// The output will display the rows as separate arrays.
	fmt.Println(matrix)

	// Using pointer to refer to the original array.
	// Example:
	// Using a pointer to refer to and modify the original array.
	// Here, we declare an array, originalArray1, with three integers: 1, 2, and 3.
	originalArray1 := [3]int{1, 2, 3}

	// We then declare a pointer variable, copiedArray1, that can point to an array of three integers.
	var copiedArray1 *[3]int

	// Assign the address of originalArray1 to the pointer copiedArray1.
	// Now, copiedArray1 points to originalArray1 in memory.
	copiedArray1 = &originalArray1

	// Change the value of the first element (index 0) of the array via the pointer.
	// Since copiedArray1 points to originalArray1, this change will affect originalArray1 as well.
	copiedArray1[0] = 100

	// Print both the original array and the dereferenced pointer.
	// Both outputs will reflect the change because they refer to the same underlying array.
	fmt.Println("Original Array: ", originalArray1)
	fmt.Println("Copied Array: ", *copiedArray1)

}

func someFunction() (int, int) {
	// This function simply returns two integers: 1 and 2.
	return 1, 2
}
