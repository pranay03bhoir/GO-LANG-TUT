package main

import (
	"fmt"
	"slices"
)

func main() {

	// Declaring a slice of integers without initializing (it's nil, has length 0).
	var numbers []int // numbers is a slice, currently empty.
	fmt.Println("numbers (uninitialized):", numbers)
	fmt.Printf("Is numbers nil? %v\n", numbers == nil)
	fmt.Printf("Length of numbers: %d\n", len(numbers))

	// Declaring and initializing a slice of integers with values 1, 2, and 3.
	var numbersSlice = []int{1, 2, 3} // numbersSlice has length 3, with specified values.
	fmt.Println("numbersSlice:", numbersSlice)
	fmt.Printf("Length of numbersSlice: %d\n", len(numbersSlice))

	// Using short variable declaration to create and initialize another slice of integers.
	numbers2 := []int{9, 8, 7} // Another method to declare and initialize a slice.
	fmt.Println("numbers2:", numbers2)
	fmt.Printf("Length of numbers2: %d\n", len(numbers2))

	// Creating a slice of integers with length 5.
	// The 'make' function allocates an underlying array and returns a slice referencing it.
	// All values are set to the zero value of int (i.e., 0).
	slice := make([]int, 5)
	fmt.Println("slice created with make:", slice)
	fmt.Printf("Length of slice: %d\n", len(slice))
	fmt.Printf("Capacity of slice: %d\n", cap(slice))

	// Declaring an array of 5 integers.
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array a:", a)

	// Slicing the array 'a' from index 1 (inclusive) to 4 (exclusive).
	// This creates a slice containing the second, third, and fourth elements of 'a'.
	slice1 := a[1:4] // slice1 will be []int{2, 3, 4}
	fmt.Println("slice1 (a[1:4]):", slice1)

	// Printing the contents of slice1 to the console.
	fmt.Println(slice1)

	// Using append to add elements to an existing slice.
	// append returns a new slice with the provided elements (6 and 7) added at the end.
	// Here, slice1 initially contains [2, 3, 4], and after this operation, it becomes [2, 3, 4, 6, 7].
	slice1 = append(slice1, 6, 7)
	fmt.Println("Appending elements to the slice1", slice1)

	// Copying a slice to another slice.
	// First, create a new slice (sliceCopy) with the same length as slice1, initialized with zero values.
	// The copy function copies the contents from slice1 to sliceCopy, element by element.
	// After copying, both slices have the same contents, but they are independent in memory.
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("This is a copy of slice1 into sliceCopy ", sliceCopy)

	// Declaring a nil slice.
	// nilSlice is declared but not initialized, so its value is nil (it has no underlying array).
	var nilSlice []int
	fmt.Println(nilSlice) // Prints: []

	// Iterating over a slice using a for loop with the range keyword.
	// The for loop gives both the index (i) and the value (v) of each element in slice1.
	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	// Accessing and modifying elements in a slice.
	// Here, we access the element at index 3 of slice1.
	fmt.Println("Element at index 3 of slice1 : ", slice1[3])

	// Modifying the element at index 3 of slice1 by assigning a new value (50).
	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1 : ", slice1[3])

	// Using Equal() utility function to compares two slices

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	// Two dimensional slices

	// twoDSlices := [][]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }

	// Creating a two-dimensional slice (a slice of slices) in Go.
	// Here, we want to create a 3x3 "matrix", which is a slice containing 3 slices,
	// each itself a slice of integers. This allows us to represent structures like grids or tables.

	// First, allocate the outer slice with length 3. Initially, each element is nil.
	twoDSlice := make([][]int, 3)

	// Now iterate over each element of the outer slice. For each index 'i':
	for i := 0; i < len(twoDSlice); i++ {
		// Determine the length of the inner slice. Here, we're setting each inner slice to also have length 3,
		// making this a 3x3 structure. (You could make them different lengths for a "jagged" slice.)
		innerLen := len(twoDSlice)

		// Allocate a new inner slice (a row) of length 3 and assign it to the outer slice at index i.
		twoDSlice[i] = make([]int, innerLen)

		// Now fill in the elements of the inner slice.
		// For each position 'j' in the inner slice:
		for j := 0; j < innerLen; j++ {
			// Set the value at position [i][j] to be i + j.
			// This means the value in each cell is the sum of its row and column indices.
			twoDSlice[i][j] = i + j
		}
	}

	// Print the resulting 2D slice.
	// The output will be: [[0 1 2] [1 2 3] [2 3 4]]
	// Each inner slice (row) shows the result of i+j for that row and its columns.
	fmt.Println(twoDSlice)

	// Creating a new slice, slice2, by slicing slice1 from index 2 (inclusive) up to index 4 (exclusive).
	// This means slice2 will include slice1[2] and slice1[3], but not slice1[4].
	slice2 := slice1[2:4]

	// Printing the contents of slice2.
	// This will display the elements at indices 2 and 3 of slice1.
	fmt.Println("Elements of slice2 (slice1[2:4]):", slice2)

	// Explaining and printing the capacity of slice2.
	// The capacity of a slice in Go is the number of elements in the underlying array, 
	// starting from the first element of the slice to the end of the underlying array.
	// For example, if slice1 has length 5, and slice2 := slice1[2:4],
	// then the capacity of slice2 will be len(slice1) - 2 = 3,
	// because slice2 starts at index 2 of slice1.
	fmt.Println("The capacity of slice2 is :", cap(slice2)) // Number of elements from slice1[2] to the end of slice1

	// Explaining and printing the length of slice2.
	// The length of a slice is simply the number of elements it holds,
	// which is equal to the difference between the high and low indices in the slicing expression.
	// In this case, slice2 := slice1[2:4], so the length is 4 - 2 = 2.
	fmt.Println("The length of slice2 is :", len(slice2))
}
