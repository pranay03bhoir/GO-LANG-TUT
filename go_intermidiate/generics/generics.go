package main

import (
	"fmt"
)

// swap is a generic function that swaps two values of any type T.
// T is a type parameter, so swap works for all types (int, string, custom types, etc).
// It returns the values in reversed order.
func swap[T any](a, b T) (T, T) {
	return b, a
}

// Stack is a generic stack data structure. It works for any type T.
type Stack[T any] struct {
	elements []T // underlying slice to hold stack elements
}

// push adds an element of type T to the top of the stack.
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
	fmt.Printf("[PUSH] Added '%v' to stack. New stack size: %d\n", element, len(s.elements))
}

// pop removes and returns the top element from the stack.
// The bool return value indicates success (true) or failure (false) if the stack is empty.
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // zero value for type T to return when stack is empty
		fmt.Println("[POP] Stack is empty, cannot pop.")
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	fmt.Printf("[POP] Removed '%v' from stack. New stack size: %d\n", element, len(s.elements))
	return element, true
}

// isEmpty returns true if the stack has no elements, false otherwise.
func (s *Stack[T]) isEmpty() bool {
	empty := len(s.elements) == 0
	fmt.Printf("[EMPTY CHECK] Is stack empty? %v\n", empty)
	return empty
}

// printAll prints all elements of the stack from bottom to top.
// If the stack is empty, it prints an appropriate message.
func (s Stack[T]) printAll() {
	fmt.Println("====== STACK CONTENTS ======")
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty.")
		fmt.Println("============================")
		return
	}
	fmt.Println("Stack elements (bottom to top):")
	for i, elem := range s.elements {
		fmt.Printf("  [%d]: %v\n", i, elem)
	}
	fmt.Println("============================")
}

func main() {

	fmt.Println("========= Generic Swap Demonstration =========")
	x, y := 1, 2
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swap:  x = %d, y = %d\n", x, y)

	x1, y1 := "Pranay", "Bhoir"
	fmt.Printf("Before swap: x1 = %s, y1 = %s\n", x1, y1)
	x1, y1 = swap(x1, y1)
	fmt.Printf("After swap:  x1 = %s, y1 = %s\n", x1, y1)
	fmt.Printf("==============================================\n")

	fmt.Println("========= Integer Stack Operations =========")
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(5)
	intStack.push(10)
	intStack.push(11)
	intStack.printAll()

	val, ok := intStack.pop()
	if ok {
		fmt.Printf("Popped value: %v\n", val)
	}
	intStack.printAll()

	val, ok = intStack.pop()
	if ok {
		fmt.Printf("Popped value: %v\n", val)
	}
	intStack.isEmpty()
	val, ok = intStack.pop()
	if ok {
		fmt.Printf("Popped value: %v\n", val)
	}
	intStack.isEmpty()
	val, ok = intStack.pop()
	if ok {
		fmt.Printf("Popped value: %v\n", val)
	}
	intStack.isEmpty()
	fmt.Printf("============================================\n")

	fmt.Println("========= String Stack Operations =========")
	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("from")
	stringStack.push("Pranay")
	stringStack.printAll()

	valStr, ok := stringStack.pop()
	if ok {
		fmt.Printf("Popped value: %v\n", valStr)
	}
	stringStack.printAll()

	stringStack.pop()
	stringStack.pop()
	stringStack.pop()
	stringStack.pop() // One extra pop to test empty stack
	stringStack.isEmpty()
	fmt.Println("============================================")
}
