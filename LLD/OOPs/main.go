package main

import (
	"fmt"
	"oops/encapsulation"
)

func encapsulation_demo() {
	fmt.Println("--------------------------------------------------------Encapsulation--------------------------------------------------------")
	user := encapsulation.NewUser("Alice", "alice@abc.com", "password123")

	// Reading Public Fields directly
	fmt.Println("Reading Public Fields directly without any getters")
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	// fmt.Printf("Password: %s\n", user.password) // This will not work, as password is private

	// Reading Private Fields using getters
	fmt.Println("\nReading Private Fields with getters")
	fmt.Printf("Login successful: %v\n", user.CheckPassword("wrong"))
	fmt.Printf("Login successful: %v\n", encapsulation.CheckPassword(user, "wrong123"))

	fmt.Printf("\nUser Details: %v\n", user) // Here private fields will also be shown, but that doesn't mean encapsulation is broken.

	fmt.Println()
	fmt.Println()

	// Attempting to modify public and private fields
	// This would work - public field
	fmt.Println("Updating Public Fields directly without any setters")
	user.Email = "alice_newid@abc.com"
	fmt.Printf("Email: %s\n", user.Email)
	// user.password = "newpassword" // Error: cannot access private field

	fmt.Println("\nUpdating Private Fields using setter")
	fmt.Printf("Updating Password\n")
	user.UpdatePassword("newpassword123")
	fmt.Printf("Login successful: %v\n", user.CheckPassword("newpassword123"))

	fmt.Printf("Updating Password\n")
	user.UpdatePassword("newpassword12345")
	fmt.Printf("Login successful: %v\n", encapsulation.CheckPassword(user, "newpassword12345"))

	fmt.Printf("\nUser Details: %v\n", user)

	fmt.Println()
	fmt.Println("Why is this called encapsulation? Because it hides the password and allows controlled access to sensitive data through getters and setters, while still allowing direct access to public fields")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")

}

func main() {
	encapsulation_demo()
	polymorphism()
	abstraction()
	inheritance()
}
