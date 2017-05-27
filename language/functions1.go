// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func generateUser() (*user, error) /* (pointer return arg, error return arg) */ {

	// Create a value of type user and return the proper values.
	someUser := user{
		name:  "hello",
		email: "hello@example.com",
		age:   25,
	}
	return &someUser, nil
}

func main() {

	// Use the function to create a value of type user. Check
	// the error being returned.
	myUser, err := generateUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value that the pointer points to.
	fmt.Println(*myUser)

	// Call the function again and just check the error.
	_, err = generateUser()
	if err != nil {
		return
	}
}
