package main

// Constants declared globally without explicitly declaring the datatype
// The go compiler automatically infers the type.
// These are called un-typed constants because we did not,
// explicitly declare there types.
const PI = 3.14
const GRAVITY = 9.81

func main() {
	// Constant declared locally with the explicitly naming the datatype.
	// These are typed constants because we declared the variable type,
	// explicitly.
	const number int = 17

	// Declaring multiple Constants.
	const (
		monday    = 1
		tuesday   = 2
		wednesday = 4
		thursday  = "thursday"
	)

}
