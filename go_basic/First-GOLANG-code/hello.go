package main // 1. THE PACKAGE DECLARATION
// "package" is a keyword that defines a directory for code organization.
// Naming this package "main" is special. It tells the Go compiler that this
// file should compile into an executable program rather than a shared library.

import "fmt" // 2. THE IMPORT STATEMENT
// "import" brings in code from other packages so you can use it.
// "fmt" (short for format) is a standard library package included with Go.
// We need it here because it contains the functionality to print text to the screen.

func main() { // 3. THE MAIN FUNCTION
// "func" is used to declare a function (a block of logic).
// "main" is a special function name. When you run a Go program, the computer
// looks for "func main()" to know where to start execution.
// The brackets "{ }" mark the beginning and end of the function's body.

    fmt.Println("Hello World") // 4. THE STATEMENT
    // Here we use the package we imported earlier ("fmt").
    // We call the "Println" (Print Line) function inside that package.
    // We pass it the string "Hello World" as an argument.
    // This sends the text to the Standard Output (your terminal/console).

} // End of the main function.