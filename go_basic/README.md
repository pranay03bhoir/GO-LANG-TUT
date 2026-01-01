# Go Language Basics - Comprehensive Tutorial

This directory contains a comprehensive collection of Go programming examples covering fundamental concepts from basic syntax to advanced features. Each subdirectory focuses on a specific topic with well-commented code examples.

## 📚 Table of Contents

1. [Getting Started](#getting-started)
2. [Variables and Constants](#variables-and-constants)
3. [Data Types and Operations](#data-types-and-operations)
4. [Control Flow](#control-flow)
5. [Loops](#loops)
6. [Collections](#collections)
7. [Functions](#functions)
8. [Error Handling and Program Flow](#error-handling-and-program-flow)
9. [Special Functions](#special-functions)

---

## Getting Started

### 1. First-GOLANG-code (`First-GOLANG-code/`)
**File:** `hello.go`

Your first Go program! This example demonstrates:
- Package declaration (`package main`)
- Import statements (`import "fmt"`)
- The `main()` function as the entry point
- Using `fmt.Println()` to print output

**Key Concepts:**
- `package main` makes the file an executable program
- `func main()` is where program execution begins
- `fmt` package provides formatted I/O functions

### 2. Import-in-go (`Import-in-go/`)
**File:** `import.go`

Learn how to import and use packages from Go's standard library:
- Multiple package imports using import blocks
- Using `net/http` package for HTTP requests
- Error handling with `if err != nil`
- `defer` statement for resource cleanup

**Key Concepts:**
- Import syntax: single and grouped imports
- Standard library packages (`fmt`, `net/http`)
- Error handling patterns in Go

---

## Variables and Constants

### 3. Variables-in-go (`variables-in-go/`)
**File:** `variables.go`

Comprehensive guide to variable declaration and usage:
- Variable declaration with `var` keyword
- Type inference with `:=` (short variable declaration)
- Global vs local scope
- Default zero values for different types
- Variable mutability

**Key Concepts:**
- `var name type = value` - explicit declaration
- `name := value` - short declaration with type inference
- Zero values: `0` for numbers, `""` for strings, `false` for bool, `nil` for pointers/slices/maps

### 4. Constants (`constants/`)
**File:** `constants.go`

Understanding constants in Go:
- Typed and untyped constants
- Global and local constant declarations
- Multiple constant declarations using blocks
- Constant immutability

**Key Concepts:**
- `const` keyword for immutable values
- Untyped constants (type inferred) vs typed constants
- Constants can be declared in blocks

---

## Data Types and Operations

### 5. Data-types-in-go (`Data-types-in-go/`)
**File:** `data-types.go`

Basic data type operations:
- String concatenation
- Numeric operations
- Boolean operations (`&&`, `||`, `!`)

**Key Concepts:**
- String concatenation with `+`
- Arithmetic operations
- Logical operators

### 6. Arithmatic-operations (`Arithmatic-operations/`)
**File:** `arithmatic_operations.go`

Comprehensive arithmetic operations:
- Addition, subtraction, multiplication, division, modulo
- Integer overflow behavior
- Floating-point operations
- Underflow examples

**Key Concepts:**
- All basic arithmetic operators (`+`, `-`, `*`, `/`, `%`)
- Integer overflow with signed and unsigned integers
- Floating-point precision and underflow

---

## Control Flow

### 7. Conditionals-in-go (`conditionals_in_go/`)
**File:** `if_else.go`

Conditional statements and branching:
- Basic `if-else` statements
- `else-if` chaining
- Nested conditionals
- User input with `fmt.Scanln()`

**Key Concepts:**
- `if condition { } else { }` syntax
- Multiple conditions with `else if`
- Nested conditionals for complex logic
- Reading user input

### 8. Switch-statements (`switch_statements/`)
**File:** `switch_case.go`

Switch statements and type switching:
- Basic switch with single values
- Multiple conditions in cases
- Switch without expression (like if-else chain)
- `fallthrough` keyword
- Type switches with `interface{}`

**Key Concepts:**
- `switch value { case x: ... }` syntax
- Multiple values in a case: `case "a", "b", "c":`
- Switch without expression: `switch { case condition: }`
- `fallthrough` to continue to next case
- Type switches: `switch x.(type) { case int: ... }`

---

## Loops

### 9. Loops (`loops/`)
**File:** `for_loop.go`

Comprehensive for loop examples:
- Traditional for loop: `for init; condition; post { }`
- Range-based iteration over slices
- `continue` and `break` statements
- Nested loops (pattern printing)
- New Go 1.22+ range over integers

**Key Concepts:**
- `for i := 0; i < n; i++` - traditional loop
- `for index, value := range collection` - range iteration
- `continue` - skip to next iteration
- `break` - exit loop early
- `for i := range 10` - iterate over integer range (Go 1.22+)

### 10. For-loop-as-while-loop (`for_loop_as_while_loop/`)
**File:** `for_as_while.go`

Using `for` loop to simulate while loop behavior:
- `for condition { }` syntax (Go doesn't have a separate `while` keyword)
- Conditional looping without initialization/post statements

**Key Concepts:**
- Go uses `for` for all loop types
- `for condition { }` acts like a while loop
- `for { }` creates an infinite loop

### 11. Guessing-game-using-while-loop (`guessing_game_using_while_loop/`)
**File:** `guessing_game.go`

Practical example combining multiple concepts:
- Random number generation
- Infinite loop with `for { }`
- User input handling
- Conditional logic
- Loop termination with `break`

**Key Concepts:**
- `rand.NewSource()` and `rand.New()` for random numbers
- `time.Now().UnixNano()` for seeding
- Infinite loops with conditional breaks
- Interactive programs

---

## Collections

### 12. Arrays (`arrays/`)
**File:** `arrays.go`

Complete guide to arrays in Go:
- Array declaration and initialization
- Array indexing and modification
- Array copying (value semantics)
- Iteration with for loops and range
- Array length with `len()`
- Array comparison
- Multidimensional arrays
- Array pointers

**Key Concepts:**
- `var arr [5]int` - fixed-size array
- `arr := [4]string{"a", "b", "c", "d"}` - initialized array
- Arrays are value types (copied, not referenced)
- `len(array)` - get array length
- Arrays can be compared with `==`
- `[3][3]int` - multidimensional arrays
- Pointers to arrays: `*[3]int`

### 13. Slices (`slices/`)
**File:** `slices.go`

Dynamic arrays (slices) in Go:
- Slice declaration and initialization
- Creating slices with `make()`
- Slicing arrays: `array[start:end]`
- Appending elements with `append()`
- Copying slices
- Nil slices
- Length and capacity
- Two-dimensional slices
- Slice comparison with `slices.Equal()`

**Key Concepts:**
- `var slice []int` - nil slice
- `slice := []int{1, 2, 3}` - slice literal
- `make([]int, length, capacity)` - create slice with make
- `array[1:4]` - slice from array (reference, not copy)
- `append(slice, elements...)` - add elements
- `copy(dest, src)` - copy slice contents
- `len(slice)` and `cap(slice)` - length and capacity
- Slices are reference types

### 14. Maps (`maps/`)
**File:** `maps.go`

Key-value pairs (maps/dictionaries):
- Creating maps with `make()`
- Adding and updating key-value pairs
- Accessing values
- Deleting keys with `delete()`
- Checking key existence (two-value assignment)
- Clearing maps with `clear()`
- Map literals
- Map comparison with `maps.Equal()`
- Iterating over maps
- Nil maps and initialization
- Nested maps

**Key Concepts:**
- `make(map[keyType]valueType)` - create map
- `map[key] = value` - add/update
- `value, exists := map[key]` - check existence
- `delete(map, key)` - remove key
- `clear(map)` - remove all keys
- Maps are reference types
- Nil maps cannot be written to
- `maps.Equal()` for comparison (Go 1.21+)

### 15. Range (`range/`)
**File:** `range.go`

The `range` keyword for iteration:
- Iterating over strings (rune by rune)
- Understanding byte positions vs rune values
- Unicode character handling

**Key Concepts:**
- `for index, rune := range string` - iterate over string
- `range` iterates over Unicode code points (runes), not bytes
- Index is byte position, value is rune (Unicode code point)
- Handles multi-byte characters correctly

---

## Functions

### 16. Functions (`functions/`)
**File:** `function.go`

Function fundamentals and advanced patterns:
- Named functions
- Anonymous functions
- Functions as first-class citizens
- Assigning functions to variables
- Passing functions as arguments
- Returning functions (higher-order functions)
- Function closures

**Key Concepts:**
- `func name(params) returnType { }` - function declaration
- Anonymous functions: `func() { }()`
- Functions are first-class: can be assigned, passed, returned
- Higher-order functions: functions that take/return functions
- Closures: functions that capture variables from outer scope

### 17. Multiple-return-value-functions (`multiple_return_value_functions/`)
**File:** `multipleReturnValueFunctions.go`

Go's powerful multiple return values:
- Functions returning multiple values
- Named return values
- Error handling pattern (value, error)
- Ignoring return values with `_`

**Key Concepts:**
- `func name() (type1, type2) { }` - multiple returns
- Named returns: `func name() (result int, err error) { }`
- `return` without values uses named return values
- `_, value := function()` - ignore first return value
- Common pattern: `(result, error)` for error handling

### 18. Variadic-functions (`variadic_functions/`)
**File:** `variadic_function.go`

Functions with variable number of arguments:
- Variadic parameter syntax (`...type`)
- Calling variadic functions
- Passing slices to variadic functions
- Multiple parameters with variadic last

**Key Concepts:**
- `func name(args ...int)` - variadic parameter
- Variadic parameter is a slice inside function
- `function(slice...)` - spread slice as arguments
- Variadic parameter must be last parameter

---

## Error Handling and Program Flow

### 19. Defer (`defer/`)
**File:** `defer.go`

Deferred function execution:
- `defer` statement basics
- LIFO (Last In, First Out) execution order
- Parameter evaluation at defer time
- Common use cases (cleanup, closing resources)

**Key Concepts:**
- `defer function()` - schedule execution after function returns
- Deferred calls execute in reverse order (stack)
- Arguments evaluated when defer is called, not when executed
- Useful for cleanup operations

### 20. Panic (`panic/`)
**File:** `panic.go`

Runtime errors and panic:
- Triggering panics with `panic()`
- Panic behavior and program termination
- Deferred functions execution during panic
- When to use panic vs error

**Key Concepts:**
- `panic(message)` - cause runtime error
- Panic stops normal execution
- Deferred functions still execute before termination
- Panic propagates up call stack
- Use for unrecoverable errors

### 21. Recover (`recover/`)
**File:** `recover.go`

Recovering from panics:
- `recover()` function usage
- Recover must be in deferred function
- Graceful error handling
- Preventing program termination

**Key Concepts:**
- `recover()` - catch and handle panics
- Must be called inside deferred function
- Returns panic value or `nil`
- Allows program to continue after panic
- Pattern: `defer func() { if r := recover(); r != nil { ... } }()`

### 22. Exit (`exit/`)
**File:** `exit.go`

Program termination with `os.Exit()`:
- Immediate program termination
- Exit codes
- Difference from panic (no deferred execution)
- When to use `os.Exit()`

**Key Concepts:**
- `os.Exit(code)` - terminate immediately
- Exit code: `0` = success, non-zero = error
- Does NOT execute deferred functions
- Use for immediate termination (different from panic)

---

## Special Functions

### 23. Init-function (`init_function/`)
**File:** `init.go`

Package initialization:
- `init()` function purpose
- Multiple `init()` functions
- Execution order
- Package initialization sequence

**Key Concepts:**
- `init()` - special function for package initialization
- Automatically called before `main()`
- Multiple `init()` functions execute in order
- Used for package setup, initialization
- Called when package is imported

---

## 🎯 Learning Path Recommendation

For beginners, follow this suggested order:

1. **Basics:**
   - First-GOLANG-code → Import-in-go → variables-in-go → constants

2. **Operations:**
   - Data-types-in-go → Arithmatic-operations

3. **Control Flow:**
   - conditionals_in_go → switch_statements

4. **Loops:**
   - loops → for_loop_as_while_loop → guessing_game_using_while_loop

5. **Collections:**
   - arrays → slices → maps → range

6. **Functions:**
   - functions → multiple_return_value_functions → variadic_functions

7. **Advanced:**
   - defer → panic → recover → exit → init_function

---

## 🚀 Running the Examples

To run any example, navigate to its directory and execute:

```bash
cd go_basic/<directory-name>
go run <filename>.go
```

For example:
```bash
cd go_basic/First-GOLANG-code
go run hello.go
```

---

## 📝 Notes

- All examples include detailed comments explaining concepts
- Code follows Go best practices and conventions
- Examples are self-contained and can be run independently
- Some examples require user input (use `fmt.Scanln()`)
- Error handling patterns are demonstrated throughout

---

## 🔗 Additional Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Tour](https://go.dev/tour/)

---

**Happy Learning! 🎉**

