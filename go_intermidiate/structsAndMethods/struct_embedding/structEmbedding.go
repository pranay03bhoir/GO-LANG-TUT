package main

import (
	"fmt"
)

// The 'person' struct represents a basic person with a name and age.
// These fields are defined using Go's struct field syntax.
type person struct {
	name string
	age  int
}

// The 'Employee' struct embeds 'person' as an anonymous field, and also adds
// extra information: empId and salary. By embedding 'person', all exported
// fields and methods of 'person' become accessible via Employee directly.
type Employee struct {
	person // Anonymous field (embedded struct)
	empId  string
	salary float64
}

// The 'introduce' method is attached to 'person'.
// It prints a greeting message including the name and age.
func (p person) introduce() {
	fmt.Printf("[person.introduce()] Hi, I am %s and I'm %d years old.\n", p.name, p.age)
}

// The 'introduce' method is also attached to 'Employee'.
// It "overrides" the 'person' introduce method for Employee values.
// It has access to directly promoted 'name' field from 'person' due to struct embedding.
// It prints employee-specific details, including promoted fields.
func (e Employee) introduce() {
	fmt.Printf("[Employee.introduce()] Hi, my name is %s, my employee ID is: %s and I earn $%.2f per year.\n", e.name, e.empId, e.salary)
}

func main() {
	// Creating an instance of Employee.
	// The 'person' part is supplied via a struct literal.
	emp := Employee{
		person: person{
			name: "Pranay",
			age:  24,
		},
		empId:  "E001",
		salary: 50000,
	}

	// ================= Detailed Output with Explanation =========================
	fmt.Println("----- Employee Details -----")
	// Accessing fields of embedded struct directly (Field Promotion)
	fmt.Printf("Accessing promoted field 'name' directly: %s\n", emp.name)
	fmt.Printf("Accessing promoted field 'age' directly: %d\n", emp.age)

	// Alternatively, we can access them using emp.person.name and emp.person.age
	fmt.Printf("Accessing with explicit path (emp.person.name): %s\n", emp.person.name)
	fmt.Printf("Accessing with explicit path (emp.person.age): %d\n", emp.person.age)

	// Accessing Employee-specific fields
	fmt.Printf("Employee ID: %s\n", emp.empId)
	fmt.Printf("Employee Salary: $%.2f\n", emp.salary)

	// Printing a separator
	fmt.Println("----- Method Demonstration -----")

	// Calling the introduce method of Employee.
	fmt.Println("Calling emp.introduce() (which uses Employee's introduce):")
	emp.introduce() // Calls Employee.introduce because it has precedence.

	// If we want, we can explicitly call the person's introduce method too:
	fmt.Println("Calling emp.person.introduce() (to explicitly use person's introduce):")
	emp.person.introduce()
}
