package main

import (
	"fmt"
)

// Shape struct demonstrates embedding in Go.
// By embedding Rectangle, Shape automatically gains access to Rectangle's methods.
type Shape struct {
	Rectangle // embedded Rectangle struct
}

// Rectangle struct defines a rectangle by its length and width.
type Rectangle struct {
	length float64 // length of the rectangle
	width  float64 // width of the rectangle
}

// Area calculates and returns the area of the rectangle.
// The receiver is Rectangle (value receiver), so this method doesn't modify the original Rectangle.
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Scale scales the dimensions of the rectangle (length and width) by the given factor.
// It uses a pointer receiver so that it can modify the original Rectangle's values.
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // Multiply length by factor
	r.width *= factor  // Multiply width by factor
}

// MyInt is a new type based on int, demonstrating method binding to custom types.
type MyInt int

// isPositive checks whether the integer is greater than zero.
// It returns true if the number is positive, false otherwise.
func (m MyInt) isPositive() bool {
	return m > 0
}

// welcomeMessage returns a simple string message.
// This method demonstrates that you can define methods without accessing the receiver.
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

func main() {

	// 1. Demonstrate Rectangle methods
	rect := Rectangle{
		length: 14,
		width:  12,
	}
	fmt.Println("Initial Rectangle dimensions:")
	fmt.Printf("Length: %.2f, Width: %.2f\n", rect.length, rect.width)

	area := rect.Area()
	fmt.Printf("Area of rectangle (%.2fx%.2f): %.2f\n", rect.length, rect.width, area)

	// Scale the rectangle and show the changes
	scaleFactor := 2.0
	rect.Scale(scaleFactor)
	fmt.Println("\nAfter scaling rectangle by factor", scaleFactor)
	fmt.Printf("New Length: %.2f, New Width: %.2f\n", rect.length, rect.width)
	area = rect.Area()
	fmt.Printf("New area of rectangle (%.2fx%.2f): %.2f\n", rect.length, rect.width, area)

	// 2. Demonstrate MyInt methods
	num := MyInt(-5)
	num1 := MyInt(9)

	fmt.Printf("\nTesting MyInt values:\n")
	fmt.Printf("Value: %d, isPositive: %v\n", num, num.isPositive())
	fmt.Printf("Value: %d, isPositive: %v\n", num1, num1.isPositive())
	fmt.Println("Welcome message from MyInt:", num.welcomeMessage())

	// 3. Demonstrate struct embedding and method expressions
	s := Shape{
		Rectangle: Rectangle{
			length: 6,
			width:  8,
		},
	}
	fmt.Printf("\nShape (embedded Rectangle) dimensions: Length %.2f, Width %.2f\n", s.length, s.width)
	// Shape has access to Rectangle methods due to embedding
	shapeArea := s.Area()
	fmt.Printf("Area of embedded Shape: %.2f\n", shapeArea)
}
