// bufio.go
// This program demonstrates how to use the bufio package in Go for buffered reading and writing.
// Each operation and its purpose are explained in detail for your understanding.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// --- BUFFERED READING DEMO ---

	// Create a new bufio.Reader that reads from a string source (implements io.Reader).
	// This is just for demonstration. In real-world, you would often wrap this around os.Stdin or a file.
	source := "Hello, bufio packageee!\nHow are you"
	reader := bufio.NewReader(strings.NewReader(source))
	fmt.Println("=== Buffered Reader Demo ===")

	// Read up to 20 bytes from the reader into the data slice.
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Printf("[!] Error occurred while reading bytes: %v\n", err)
		return
	}
	// Display exactly what was read, as a quoted string for clarity.
	fmt.Printf("[*] Read %d bytes: %q\n", n, data[:n])

	// Continue reading, but this time until the next newline character is seen (including '\n').
	// This is useful for line-based input, such as reading from a terminal or file line by line.
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("[!] Error occurred while reading until newline: %v\n", err)
		return
	}
	fmt.Printf("[*] Read line (up to next newline): %q\n", line)

	// --- BUFFERED WRITING DEMO ---

	// Create a new bufio.Writer that buffers output before writing to the final destination (here: os.Stdout).
	writer := bufio.NewWriter(os.Stdout)
	fmt.Println("=== Buffered Writer Demo ===")

	// Write some bytes to buffer (not printed yet!).
	dataToWrite := []byte("Hello, bufio package!\n")
	nn, err := writer.Write(dataToWrite)
	if err != nil {
		fmt.Printf("[!] Error occurred while writing bytes: %v\n", err)
		return
	}
	fmt.Printf("[*] Buffered %d bytes for output: %q\n", nn, dataToWrite)

	// The above write doesn't actually display the output yet; it's in the buffer.
	// We need to call Flush() to ensure all buffered data is written out!
	err = writer.Flush()
	if err != nil {
		fmt.Printf("[!] Error occurred while flushing buffer: %v\n", err)
		return
	}
	fmt.Println("[*] Flushed writer buffer to output (should see line above).")

	// Write a string (with automatic conversion to bytes and buffering).
	sampleStr := "This is a String\n"
	n, err = writer.WriteString(sampleStr)
	if err != nil {
		fmt.Printf("[!] Error occurred while writing string: %v\n", err)
		return
	}
	fmt.Printf("[*] Buffered %d bytes for output from string: %q\n", n, sampleStr)

	// Flush again to ensure the buffered string is actually sent to the output stream.
	err = writer.Flush()
	if err != nil {
		fmt.Printf("[!] Error occurred while flushing writer after string: %v\n", err)
		return
	}
	fmt.Println("[*] Flushed writer buffer to output (should see string above).")

	// --- SUMMARY OF WHAT YOU LEARNED ---
	// 1. bufio.Reader can wrap any io.Reader to provide buffered, efficient reading
	//    - Use Read to grab up to N bytes
	//    - Use ReadString('\n') to read up to (and including) the next newline
	// 2. bufio.Writer wraps any io.Writer to provide buffered writing
	//    - Write and WriteString add to the buffer
	//    - Use Flush() to ensure your data is pushed out (important when writing to files or network!)
	// 3. Buffering can greatly improve performance when dealing with large data or slow destinations.
}
