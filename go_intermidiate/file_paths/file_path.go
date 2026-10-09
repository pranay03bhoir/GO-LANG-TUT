// This program demonstrates several essential file path manipulations in Go using the "path/filepath" and "strings" packages.
// Detailed explanations and improved print statements are provided for enhanced clarity.

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Define both a relative and an absolute file path as string variables
	relativePath := "./Documents/file.txt"                    // A relative path (relative to the current directory)
	absolutePath := "/home/user/Documents/file.txt"           // An absolute file path (starts from the root)

	// Join multiple path segments into a single well-formed path
	joinedPath := filepath.Join("home", "user", "downloads", "file.zip")
	fmt.Printf("[Join Example] Input Segments: 'home', 'user', 'downloads', 'file.zip'\n")
	fmt.Printf("Joined Path: %q\n\n", joinedPath)

	// Clean a file path by simplifying it, removing unnecessary navigation such as ./ and ../
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Printf("[Clean Example] Original: %q\n", "./data/../data/file.txt")
	fmt.Printf("Normalized (Cleaned) Path: %q\n\n", normalizedPath)

	// Split a path into its directory and file parts
	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Printf("[Split Example]\n")
	fmt.Printf("Input: %q\n", "/home/user/docs/file.txt")
	fmt.Printf("Directory: %q\n", dir)
	fmt.Printf("File: %q\n\n", file)

	// Extract just the last element (final directory or file name) from a path
	base := filepath.Base("/home/user/docs/")
	fmt.Printf("[Base Example] Input: %q\n", "/home/user/docs/")
	fmt.Printf("Base (last element): %q\n\n", base)

	// Check whether given paths are absolute
	fmt.Printf("[IsAbs Example]\n")
	fmt.Printf("Is relativePath (%q) absolute? %v\n", relativePath, filepath.IsAbs(relativePath))
	fmt.Printf("Is absolutePath (%q) absolute? %v\n\n", absolutePath, filepath.IsAbs(absolutePath))

	// Extract file extension and strip extension for just the file name
	extension := filepath.Ext(file)
	filename := strings.TrimSuffix(file, extension)
	fmt.Printf("[Extension & Filename Example]\n")
	fmt.Printf("File: %q\n", file)
	fmt.Printf("Extension: %q\n", extension)
	fmt.Printf("File name without extension: %q\n\n", filename)

	// Find the relative path between two locations
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	fmt.Printf("[Rel Example 1] Path from base %q to target %q\n", "a/b", "a/b/t/file")
	if err != nil {
		fmt.Printf("Error determining relative path: %v\n", err)
	} else {
		fmt.Printf("Relative Path: %q\n\n", rel)
	}

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	fmt.Printf("[Rel Example 2] Path from base %q to target %q\n", "a/c", "a/b/t/file")
	if err != nil {
		fmt.Printf("Error determining relative path: %v\n", err)
	} else {
		fmt.Printf("Relative Path: %q\n\n", rel)
	}

	// Convert a relative path to an absolute path
	absPath, err := filepath.Abs(relativePath)
	fmt.Printf("[Abs Example] Converting relative path %q to absolute path\n", relativePath)
	if err != nil {
		fmt.Printf("Error obtaining absolute path: %v\n", err)
	} else {
		fmt.Printf("Absolute Path: %q\n", absPath)
	}
}
