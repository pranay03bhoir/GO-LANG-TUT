// base64_encode_decode.go
// This program demonstrates in VERY DETAILED steps how to perform standard and URL-safe Base64 encoding and decoding in Go.
// Each operation is explained thoroughly to aid your understanding and learning!

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// ---------------------------------------------------
	// 1. Original Data: Let's use a simple string for demo
	// ---------------------------------------------------
	data := []byte("Hello, Base64 Encoding")
	// Note: We convert the string to a slice of bytes because most encoding functions operate on []byte.

	fmt.Println("=== Step 1: ORIGINAL DATA ===")
	fmt.Printf("Original string: %q\n", string(data))
	fmt.Printf("Original byte values: %v\n\n", data)

	// ---------------------------------------------------
	// 2. Standard Base64 Encoding
	// ---------------------------------------------------
	// Base64 encoding converts binary data to an ASCII string format (safe for text protocols and storage).
	// Standard encoding uses A-Z, a-z, 0-9, '+', and '/'.

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("=== Step 2: STANDARD BASE64 ENCODING ===")
	fmt.Printf("Base64 encoded string: %q\n", encoded)
	fmt.Println("This string can be safely transmitted/embedded in places where binary data is not allowed.")

	// ---------------------------------------------------
	// 3. Decoding from Base64 back to the original data
	// ---------------------------------------------------
	// DecodeString takes a Base64-encoded string and returns the original bytes.
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("[!] Error decoding base64 data:", err)
		return
	}
	fmt.Println("=== Step 3: DECODE BASE64 BACK TO ORIGINAL ===")
	fmt.Printf("Decoded bytes: %v\n", decoded)
	fmt.Printf("Decoded string: %q\n\n", string(decoded))

	// ---------------------------------------------------
	// 4. URL-Safe Base64 Encoding
	// ---------------------------------------------------
	// Standard Base64 uses '+' and '/' which may not work safely in URLs or filenames.
	// Go provides URLEncoding which replaces '+' with '-' and '/' with '_'.

	data = []byte("He~lo, Base64 Encoding") // changing input to show it's safe for ~, etc.
	fmt.Println("=== Step 4: URL-SAFE BASE64 ENCODING ===")
	fmt.Printf("Original string for URL-safe encoding: %q\n", string(data))
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Printf("URL-safe Base64 encoding: %q\n", urlSafeEncoded)
	fmt.Println("This encoded string is safe to use in URLs, HTTP GET parameters, and filenames!")

	// ---------------------------------------------------
	// SUMMARY AND LEARNING POINTS
	// ---------------------------------------------------
	fmt.Println(">>> SUMMARY <<<")
	fmt.Println("1. Use base64.StdEncoding for regular data encoding/decoding.")
	fmt.Println("2. Use base64.URLEncoding for data that will go into URLs or filenames.")
	fmt.Println("3. Always handle errors when decoding, as the input string must be a valid base64 format.")
	fmt.Println("4. Base64 is an ENCODING, not encryption. Do NOT use it for securing secrets!")
	fmt.Println("5. The output is always ASCII text, regardless of the original input bytes.")
}
