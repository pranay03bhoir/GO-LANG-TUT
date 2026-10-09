package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// This Go program demonstrates how to read a text file line by line and filter lines containing
// a specific keyword. When such a line is found, it replaces the keyword with another word
// and prints both the original and updated versions with clear formatting for easy understanding.
//
// Step-by-step explanation:
// 1. The program opens a file named "example.txt" for reading.
// 2. It checks for errors when opening the file; if an error occurs, it prints an informative message and exits.
// 3. It ensures that the file will be closed properly by deferring file.Close() until the end of main().
// 4. A bufio.Scanner is used to read the file line by line, which is efficient for text files.
// 5. The keyword to search for is set to "important".
// 6. The program iterates through each line, checking if the keyword exists within the line.
// 7. If the keyword is found, it creates an updated version of the line with all instances of the keyword replaced by "necessary".
// 8. Both the original and updated lines are printed with improved and detailed print statements for clarity.
// 9. After reading the file, the code checks if any error occurred during scanning, printing an error message if so.

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file 'example.txt': %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	keyword := "important"

	lineNumber := 1 // Keep track of the line number for improved output
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Printf("-- Line %d matches filter --\n", lineNumber)
			fmt.Printf("Original: %s\n", line)
			fmt.Printf("Modified: %s\n", updatedLine)
			fmt.Println("--------------------------------------------------")
			lineNumber++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occurred while scanning 'example.txt': %v\n", err)
	}
}
