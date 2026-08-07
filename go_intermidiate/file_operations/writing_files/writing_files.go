
// Go Program: Writing Data to Files with Detailed Explanation and Improved Print Statements

/*
   This program demonstrates how to create and write data to files in Go using the "os" package.
   Every step is explained for better understanding — perfect for learning how file operations work in Go!

   We'll create two files:
     1. output.txt      — written with []byte data using file.Write
     2. writeString.txt — written with string data using file.WriteString
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// ----------- 1. Create a New File "output.txt" -----------
	fmt.Println("[Step 1] Creating a new file named 'output.txt'...")

	// os.Create attempts to create the file.
	// If the file already exists, it truncates (clears) it to zero length.
	file, err := os.Create("output.txt")
	if err != nil {
		// If error is not nil, file creation failed (e.g., permissions issue)
		fmt.Println("✗ Error: Failed to create 'output.txt':", err)
		return
	}
	// Ensures that file.Close() will be called at the end of main(), releasing the resource.
	defer file.Close()
	fmt.Println("✓ 'output.txt' created successfully.")

	// ----------- 2. Write []byte Data to "output.txt" -----------
	fmt.Println("[Step 2] Writing byte data to 'output.txt'...")

	// Data to write. \n ensures new line at the end.
	data := []byte("Hello World\n")
	// file.Write writes the byte slice to the file.
	numBytes, err := file.Write(data)
	if err != nil {
		fmt.Println("✗ Error: Failed to write to 'output.txt':", err)
		return
	}
	fmt.Printf("✓ Wrote %d bytes to 'output.txt': %q\n", numBytes, data)

	// ----------- 3. Create a New File "writeString.txt" -----------
	fmt.Println("\n[Step 3] Creating a new file named 'writeString.txt'...")

	// If you reuse variable names (like file), previous file descriptor should be closed (handled by defer).
	file2, err := os.Create("writeString.txt")
	if err != nil {
		fmt.Println("✗ Error: Failed to create 'writeString.txt':", err)
		return
	}
	// Defer closing of this second file.
	defer file2.Close()
	fmt.Println("✓ 'writeString.txt' created successfully.")

	// ----------- 4. Write String Data to "writeString.txt" -----------
	fmt.Println("[Step 4] Writing string data to 'writeString.txt'...")

	numChars, err := file2.WriteString("Hello Go !!!\n")
	if err != nil {
		fmt.Println("✗ Error: Failed to write to 'writeString.txt':", err)
		return
	}
	fmt.Printf("✓ Wrote %d bytes to 'writeString.txt': %q\n", numChars, "Hello Go !!!\n")

	// ----------- 5. Finished -----------
	fmt.Println("\nAll file writing operations completed successfully!")
}

/*
   --- DETAILED LEARNING POINTS ---

   1. os.Create(filename):
        - Creates a new file for writing. If the file exists, its length is truncated to zero.
        - Returns an *os.File (the file handle) and an error if any.

   2. defer file.Close():
        - Schedules file.Close() to run when the function (main) finishes, ensuring resources are freed.
        - Good practice: always close file descriptors when done.

   3. file.Write([]byte):
        - Writes raw byte data to the file.
        - Returns number of bytes written and an error.

   4. file.WriteString(string):
        - Writes a string to the file (convenient for simple text).
        - Writes and returns number of bytes (not runes/characters).

   5. Error Checking:
        - Always check for errors after file operations, as IO can fail due to permissions, disk, etc.
*/