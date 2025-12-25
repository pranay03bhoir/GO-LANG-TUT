package main

import (
	"fmt"
	"maps"
)

func main() {
	// Step 1: Creating a map in Go
	// This statement creates a variable called myMap, which is a map where the keys are strings and the values are integers.
	// The make() function initializes the map so it's ready for use, but it's empty at this point.
	myMap := make(map[string]int)
	fmt.Println("Initial map (empty):", myMap) // Prints: Initial map (empty): map[]

	// Step 2: Adding entries to the map
	// You add items (key-value pairs) to a map using the syntax: map[key] = value
	// Here, we're setting "key1" to 9, and "code" to 18.
	myMap["key1"] = 9  // Adds a new entry: "key1" : 9
	myMap["code"] = 18 // Adds a new entry: "code" : 18

	fmt.Println("Map after adding key1 and code:", myMap)
	// Output may look like: Map after adding key1 and code: map[code:18 key1:9]

	// Step 3: Accessing a value by key
	// You can access a value by referencing its key in the map; if the key does not exist, you get the zero value for the map's value type (here, 0).
	value1 := myMap["key1"]                               // Retrieves the value associated with "key1", which is 9.
	fmt.Printf("Accessed value for 'key1': %d\n", value1) // Prints: Accessed value for 'key1': 9

	// Step 4: Modifying a value for an existing key
	// If you assign a different value to an existing key, the value is updated.
	myMap["code"] = 22 // Updates the value for "code" from 18 to 22
	fmt.Println("Map after updating 'code' to 22:", myMap)
	// Output may look like: Map after updating 'code' to 22: map[code:22 key1:9]

	// Step 5: Deleting a key-value pair
	// The delete function removes an entry from the map. If the key doesn't exist, nothing happens.
	delete(myMap, "key1") // Removes the entry with key "key1"
	fmt.Println("Map after deleting 'key1':", myMap)
	// Output may look like: Map after deleting 'key1': map[code:22]

	// Step 6: Adding more key-value pairs
	// You can continue adding new entries to the map as needed.
	myMap["key2"] = 9
	myMap["key3"] = 10
	myMap["key4"] = 19
	myMap["key5"] = 99
	fmt.Println("Updated map with more pairs:", myMap)
	// The map now contains keys "code", "key2", "key3", "key4", "key5" with their respective values.

	// Step 7: Retrieving a value and checking for key existence
	// When accessing a map, Go allows you to use a two-value assignment: val, exists := myMap["key"]
	// - val is the value for the key (or the zero value if key not present).
	// - exists is a bool indicating if the key was actually present in the map.
	val, exists := myMap["key2"]
	fmt.Printf("Value for 'key2': %d (exists: %v)\n", val, exists) // Prints: Value for 'key2': 9 (exists: true)

	// If you try to get a value for a key that's not in the map, exists will be false and val will be zero
	nonExistentVal, nonExistentExists := myMap["notInMap"]
	fmt.Printf("Value for 'notInMap': %d (exists: %v)\n", nonExistentVal, nonExistentExists) // Prints: Value for 'notInMap': 0 (exists: false)

	// This line re-demonstrates two-value assignment, but here the second variable (unknownValue) would be the existence boolean
	value, unknownValue := myMap["key2"]
	fmt.Println(value)        // Prints the value for "key2" (e.g., 9)
	fmt.Println(unknownValue) // Prints true if "key2" exists; otherwise false

	// Step 8: Clearing all entries from the map
	// The clear() function removes all key-value pairs, leaving the map empty but still allocated.
	clear(myMap)
	fmt.Println("Map after clearing (should be empty):", myMap)
	// Output: Map after clearing (should be empty): map[]

	//===========================================================================//

	// Declaring and initializing two new maps: myMap2 and myMap3.
	// Both have the same key-value pairs: "a": 1 and "b": 2.
	myMap2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	myMap3 := map[string]int{
		"a": 1,
		"b": 2,
	}

	// Print the entire myMap2 to show its contents.
	fmt.Println("myMap2 contents:", myMap2)

	// Checking equality between two maps using maps.Equal from the "maps" package (Go 1.21+).
	// maps.Equal returns true if both maps contain identical key-value pairs.

	// First, compare myMap (which has been cleared earlier) to myMap2.
	// Since myMap is now empty, this should be false.
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("myMap and myMap2 are NOT equal (expected, because myMap is empty)")
	}

	// Next, compare myMap3 to myMap2.
	// These have the exact same contents, so this should print that they are equal.
	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal (both have the same key-value pairs)")
	} else {
		fmt.Println("myMap3 and myMap2 are NOT equal")
	}

	// Iterating over myMap2 to print each key and its corresponding value in detail.
	fmt.Println("Iterating over myMap2:")
	for key, value := range myMap2 {
		fmt.Printf("  Key: '%s', Value: %d\n", key, value)
	}

	// Demonstrating nil maps, key lookups, assignment, and length

	// Declare a map variable without initialization. At this point, myMap4 is nil:
	var myMap4 map[string]string

	// Check if myMap4 is nil (i.e., it doesn't point to an initialized map structure)
	if myMap4 == nil {
		fmt.Println("myMap4 is nil (uninitialized)")
	} else {
		fmt.Println("myMap4 is NOT nil (it has been initialized)")
	}

	// Try to access a key in a nil map. This is safe: Go treats missing keys as returning the zero value for the map's value type.
	// In this case, since the value type is string, the result will be "" (empty string) if "key" is not present, even though the map is nil.
	valu, exist := myMap4["key"]
	fmt.Printf("Looking up 'key' in nil map: value='%s', exists=%v\n", valu, exist)
	// Note: In Go, 'exists' is always false for missing keys, regardless of whether the map is nil or empty.

	// Attempting to assign to a nil map will panic! Uncommenting this line would cause a runtime error.
	// myMap4["firstName"] = "pranay"

	// Proper way: Allocate a new map before inserting any key-value pairs.
	myMap4 = make(map[string]string)

	// Now, map is initialized; we can safely set a value.
	myMap4["firstName"] = "pranay"

	// Print the entire map so we see its contents.
	fmt.Println("Contents of myMap4 after adding 'firstName':", myMap4)

	// Show the value for "firstName" using explicit lookup
	valueFName, existsFName := myMap4["firstName"]
	fmt.Printf("Lookup of 'firstName': value='%s', exists=%v\n", valueFName, existsFName)

	// Show that accessing a non-existent key after initialization still returns the zero value, not an error.
	valueLastName, existsLastName := myMap4["lastName"]
	fmt.Printf("Lookup of 'lastName': value='%s', exists=%v\n", valueLastName, existsLastName)

	// Print the length of the map
	fmt.Printf("myMap4 length is: %d (should be 1 after adding 'firstName')\n", len(myMap4))

	// Nested maps in go.

	// Creating a nested map in Go: a map whose values are themselves maps.
	// Here, myMap5 is of type map[string]map[string]string, meaning it maps strings to (maps from strings to strings).
	myMap5 := make(map[string]map[string]string) // Initialize the outer map.

	// Assign the previously created myMap4 (which is of type map[string]string) to myMap5 under the key "map1".
	// This demonstrates how you can use a regular map as a value inside a nested map structure.
	myMap5["map1"] = myMap4

	// Print the entire nested map to visualize the structure.
	// The output will look something like: map[map1:map[firstName:pranay]]
	// This shows that the key "map1" in myMap5 maps to another map containing the pair "firstName":"pranay".
	fmt.Println(myMap5)
}
