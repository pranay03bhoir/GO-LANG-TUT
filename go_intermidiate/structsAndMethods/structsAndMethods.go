package main

import (
	"fmt"
)

// Person struct demonstrates the use of nested struct fields, 
// both named and embedded (anonymous). 
type Person struct {
	firstName   string
	lastName    string
	age         int
	address     Address    // Named field: Address struct is explicitly named.
	PhoneNumber             // Embedded/anonymous field: methods/fields of PhoneNumber become part of Person.
}

// PhoneNumber struct holds home and cell numbers.
type PhoneNumber struct {
	homeCell  string
	cellPhone string
}

// Address struct represents a simple geographic address.
type Address struct {
	city    string
	country string
}

func main() {
	fmt.Println("==== Go Structs Detailed Demonstration ====")

	// -- 1. Creating struct variables with field names --

	// Creating a Person instance with most fields filled; address and PhoneNumber have zero/default values.
	p := Person{
		firstName: "Pranay",
		lastName:  "Bhoir",
		age:       22,
	}
	fmt.Println("\n1. A 'Person' struct instance created with minimal fields:")
	fmt.Printf("p=%+v\n", p)
	fmt.Println("-- Note: address and phone numbers are blank (zero values)")

	// Creating a Person with only first name and age
	p1 := Person{
		firstName: "John",
		age:       33,
	}
	fmt.Println("\n2. Another 'Person' with only first name and age:")
	fmt.Printf("p1=%+v\n", p1)
	fmt.Println("-- Again, lastName, address, and phone numbers are at zero values")

	// -- 2. Creating struct with nested and embedded structs filled in --
	p2 := Person{
		firstName: "Kalpana",
		lastName:  "Chawla",
		age:       45,
		address: Address{
			city:    "Mumbai",
			country: "India",
		},
		PhoneNumber: PhoneNumber{
			homeCell:  "9850676710",
			cellPhone: "8830986365",
		},
	}

	fmt.Println("\n3. A fully initialized 'Person' with nested Address and embedded PhoneNumber:")
	fmt.Printf("p2=%+v\n", p2)
	fmt.Println("-- Accessing each field for demonstration:")

	fmt.Printf("Full Name: %s %s\n", p2.firstName, p2.lastName)
	fmt.Printf("Age: %d\n", p2.age)
	fmt.Printf("Address: %s, %s\n", p2.address.city, p2.address.country)
	// Because PhoneNumber is embedded, we can access its fields directly:
	fmt.Printf("Home Phone: %s\n", p2.homeCell)
	fmt.Printf("Cell Phone: %s\n", p2.cellPhone)

	// -- 3. Modifying fields of a struct after creation --
	fmt.Println("\n4. Modifying the address fields of 'p': Changing city and country:")
	fmt.Printf("Before: p.address = %+v\n", p.address)
	p.address.city = "New York"
	p.address.country = "USA"
	fmt.Printf("After: p.address = %+v\n", p.address)

	// -- 4. Comparing fields, showing independent values of structs --
	fmt.Println("\n5. Display first names to show independent fields across structs:")
	fmt.Printf("p.firstName = %q\n", p.firstName)
	fmt.Printf("p1.firstName = %q\n", p1.firstName)
	fmt.Printf("p.address.city = %q\n", p.address.city)
	fmt.Printf("p.address.country = %q\n", p.address.country)

	// -- 5. Anonymous struct example --
	fmt.Println("\n6. Anonymous struct usage:")
	user := struct {
		userName string
		email    string
	}{
		userName: "user123",
		email:    "psuedoemail232@example.org",
	}
	fmt.Printf("Anonymous user struct: %+v\n", user)
	fmt.Printf("userName: %q, email: %q\n", user.userName, user.email)
	fmt.Println("-- Anonymous structs are convenient for ad-hoc data structures --")

	// -- 6. Methods on struct types --
	fmt.Println("\n7. Using methods on Person struct:")

	fmt.Printf("Calling p.fullName(): %q\n", p.fullName()) // Uses receiver by value

	fmt.Printf("Before increment, p.age = %d\n", p.age)
	p.incrementAgeByOne() // Uses receiver by pointer, mutates struct
	fmt.Printf("After increment, p.age = %d\n", p.age)

	// -- 7. Comparing structs --
	fmt.Println("\n8. Comparing two Person structs (p vs p1):")
	fmt.Printf("p == p1 ? %v\n", p == p1)
	fmt.Println("-- Note: Structs are comparable if all their fields are comparable, "+
		"and all field values (including embedded structs) must match for == to be true.")

	fmt.Println("\n==== End of Struct Demonstration ====")
}

// Method on Person: Returns their full name by concatenating first and last.
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Pointer receiver: Updates age field by incrementing it by 1.
func (p *Person) incrementAgeByOne() {
	p.age++
}
