
/*
   Go Program: Reading Data from a File with Detailed Explanation and Improved Print Statements

   This program demonstrates how to open a file and read its entire content, as well as how to scan and print each line using bufio.Scanner.
   Every step is commented in detail for your deep understanding and learning!
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ----------- 1. Attempt to Open the File -----------
	fmt.Println("[Step 1] Attempting to open 'output.txt' for reading...")

	/*
	   os.Open("output.txt")
	   - Opens the named file for reading only.
	   - Returns an *os.File object and an error value.
	*/
	file, err := os.Open("output.txt")
	if err != nil {
		// If there is a problem (file missing, permissions), handle it here.
		fmt.Printf("✗ Error: Failed to open 'output.txt': %v\n", err)
		return
	}
	// Use defer to ensure file.Close() executes when main() returns.
	defer func() {
		fmt.Println("[Info] Closing 'output.txt'.")
		file.Close()
	}()
	fmt.Println("✓ Successfully opened 'output.txt'.")

	// ----------- 2. Reading the Entire File Content as Bytes -----------
	fmt.Println("\n[Step 2] Reading entire content of the file as bytes...")

	/*
	   Here, we allocate a byte slice (buffer) of size 1024 bytes.
	   We read up to 1024 bytes from the file into this buffer using file.Read().
	   Note: file.Read reads from the current file cursor, which starts at the beginning.
	*/
	data := make([]byte, 1024) // Allocate a buffer
	numBytes, err := file.Read(data)
	if err != nil {
		fmt.Printf("✗ Error: Failed to read data from file: %v\n", err)
		return
	}
	fmt.Printf("✓ Read %d bytes from 'output.txt'.\n", numBytes)
	fmt.Println("[Output] File content (as string):")
	fmt.Printf("--------------------------------------------------\n%s\n--------------------------------------------------\n", string(data[:numBytes]))

	// ----------- 3. Resetting File Cursor for Line-by-Line Reading -----------
	/*
	   Since file.Read has already read to the end or advanced the file pointer,
	   bufio.Scanner will see an empty file unless we reset the file's read pointer to the start.
	   We use file.Seek(0, 0) to reset it.
	*/
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Printf("✗ Error: Failed to reset file cursor: %v\n", err)
		return
	}
	fmt.Println("\n[Step 3] Now scanning and printing each line individually:")

	// ----------- 4. Scanning and Printing Each Line -----------
	/*
	   bufio.NewScanner wraps the file and allows you to iterate through it line by line.
	   scanner.Scan() advances the scanner to the next line; scanner.Text() returns the line string.
	   Check scanner.Err() after the loop for any reading errors.
	*/
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text() // Retrieve the current line as a string
		fmt.Printf("  Line %d: %q\n", lineNumber, line)
		lineNumber++
	}
	// After looping, check for errors (other than EOF)
	if err := scanner.Err(); err != nil {
		fmt.Printf("✗ Error: Encountered problem during line-by-line reading: %v\n", err)
		return
	}
	fmt.Printf("✓ Finished reading all %d lines from 'output.txt'.\n", lineNumber-1)

	// ----------- 5. Wrap-up Message -----------
	fmt.Println("\nAll file reading operations completed successfully!")
}

/*
   --- DETAILED LEARNING POINTS ---

   1. os.Open(filename):
        - Opens the file in read-only mode.
        - Returns a file handle and an error if failure occurs.

   2. defer file.Close():
        - Ensures resources are released properly after you are done with the file.

   3. file.Read():
        - Reads up to len(buffer) bytes from file into the buffer.
        - Returns number of bytes actually read, and an error.

   4. file.Seek(offset, whence):
        - Moves the file cursor to a specific position.
        - Needed if you want to re-read the file from the beginning.

   5. bufio.NewScanner(file):
        - Reads the file line by line.
        - scanner.Scan() moves to the next line, scanner.Text() returns it as a string.
        - Always check scanner.Err() for errors after reading.

   6. Print Statements:
        - Provide detailed, user-friendly feedback for each stage of the program.
        - Use ✓ and ✗ icons for success/failure clarity.
        - Show both full content and per-line reading for learning contrast.
*/