// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import (
	"fmt"
	"math"
)

func main() {

	// Declare a nil slice of integers.
	var slice []int

	fmt.Println("Slice Length: ", len(slice))
	fmt.Println("Slice Capacity: ", cap(slice))

	// Appends numbers to the slice.
	for i := 0; i < 10; i++ {
		slice = append(slice, int(math.Pow(2, float64(i))))
	}

	// Display each value in the slice.
	for _, item := range slice {
		fmt.Println(item)
	}
	// Declare a slice of strings and populate the slice with names.
	names := []string{"Abel", "Bill", "Cathy", "Daniel", "Ernest"}

	// Display each index position and slice value.
	for i, name := range names {
		fmt.Println(i, name)
	}

	fmt.Println()

	// Take a slice of index 1 and 2 of the slice of strings.
	subslice := names[1:3]

	// Display each index position and slice values for the new slice.
	for i, name := range subslice {
		fmt.Println(i, name)
	}
}
