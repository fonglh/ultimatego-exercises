// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	names := make(map[string]int)

	// Initialize some data into the map.
	names["Abel"] = 30
	names["Bozz"] = 35
	names["Easy"] = 26
	names["Bob"] = 8
	names["Kate"] = 8

	// Display each key/value pair.
	for key, value := range names {
		fmt.Println("Key:", key, ", Value:", value)
	}
}
