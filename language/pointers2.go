// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
func changeName(oldName *string, newName string) {

	// Use the pointer to change the value that the
	// pointer points to.
	*oldName = newName
}

// Use a pointer to the full user instead.
// This might be less desirable because the function needs to know
// about the fields in the struct.
func changeEmail(userPtr *user, newEmail string) {
	userPtr.email = newEmail
}

func main() {

	// Create a variable of type user and initialize each field.
	user1 := user{name: "bill", email: "bill@example.com", age: 30}

	// Display the value of the variable.
	fmt.Println(user1.name)
	fmt.Println(user1.email)

	// Share the variable with the function you declared above.
	changeName(&user1.name, "william")
	changeEmail(&user1, "william@example.com")

	// Display the value of the variable.
	fmt.Println(user1.name)
	fmt.Println(user1.email)
}
