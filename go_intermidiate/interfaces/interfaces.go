// Detailed Explanation and Improved Printing for the Interfaces Example in Go

package main

import (
	"fmt"
	"math"
)

// geometry interface declares two methods: area and perimeter.
// Any type that implements these two methods is implicitly of type geometry.
type geometry interface {
	area() float64      // returns the area of the shape
	perimeter() float64 // returns the perimeter of the shape
}

// rectangle struct models a rectangle with width and height.
type rectangle struct {
	width, height float64
}

// rect is another struct identical in fields to rectangle,
// but is a different type. This is to demonstrate that having 
// the same structure doesn't mean the methods are shared.
type rect struct {
	width, height float64
}

// circle struct models a circle with a radius.
type circle struct {
	radius float64
}

// area calculates area for rectangle using value receiver.
// Implements the area() method from geometry interface.
func (r rectangle) area() float64 {
	return r.height * r.width
}

// area for rect (different type from rectangle), not implementing perimeter.
// This means rect does NOT implement the geometry interface fully.
func (r rect) area() float64 {
	return r.height * r.width
}

// area for circle using the formula πr²
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perimeter for rectangle using the formula 2*(width+height)
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// perimeter for circle using the formula 2πr
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// diameter returns the diameter of the circle; not part of geometry interface.
func (c circle) diameter() float64 {
	return 2 * c.radius
}

// measure accepts any geometry type and prints detailed info.
// Improved printing for clarity and understanding.
func measure(g geometry) {
	// Type switch to print the kind of geometric shape
	fmt.Println("--------------------------------------------------")
	switch shape := g.(type) {
	case rectangle:
		fmt.Printf("Shape details (Rectangle): width=%.2f, height=%.2f\n", shape.width, shape.height)
	case circle:
		fmt.Printf("Shape details (Circle): radius=%.2f\n", shape.radius)
	default:
		fmt.Printf("Shape details (Unknown type): %+v\n", g)
	}
	fmt.Printf("Area:      %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
	// Bonus: print diameter if available
	if c, ok := g.(circle); ok {
		fmt.Printf("Diameter:  %.2f\n", c.diameter())
	}
	fmt.Println("--------------------------------------------------")
}

// myPrinter prints any number of values separated by spaces.
// Adds indices and value type for demonstration.
func myPrinter(i ...interface{}) {
	fmt.Println("Custom Printer Output:")
	for idx, v := range i {
		fmt.Printf("[%d: %T] %v ", idx, v, v)
	}
	fmt.Println() // Newline for readability
}

// printType prints the type of the provided value using type switching.
func printType(i interface{}) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Input value: %v\n", i)
	switch i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case float64:
		fmt.Println("Type: float64")
	case rune:
		fmt.Println("Type: rune")
	case bool:
		fmt.Println("Type: bool")
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println("--------------------------------------------------")
}

func main() {
	// Create a rectangle instance
	r := rectangle{
		width:  3,
		height: 5,
	}

	// Alternative rect type instance for demonstration (uncomment to observe interface failure)
	// r1 := rect{
	// 	width:  3,
	// 	height: 5,
	// }

	// Create a circle instance
	c := circle{
		radius: 5,
	}

	fmt.Println("Measuring rectangle:")
	measure(r) // rectangle implements geometry

	fmt.Println("Measuring circle:")
	measure(c) // circle implements geometry

	// Uncommenting below will cause compile error since rect doesn't implement perimeter.
	// measure(r1) // rect does not satisfy geometry interface

	// Demonstrating custom printer with mixed types (including boolean, string, float, etc)
	myPrinter(1, "pear", "B", 77.89, true)

	// Demonstrating dynamic type printing, including types not explicitly handled by printType
	printType(false)   // bool
	printType(99.99)   // float64
	printType("PB")    // string
	printType('s')     // rune (unicode code point for 's')
	printType([]int{1,2,3}) // Unknown type

	fmt.Println("Done with demonstration!")
}
