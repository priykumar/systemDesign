package main

import "fmt"

type Authenticator interface {
	CheckPassword(password string) bool
	UpdatePassword(newPassword string)
}

type User struct {
	Name     string // Public fields
	Email    string // Public fields
	password string // Private fields
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		password: password,
	}
}

// Getters for private fields
func (u *User) CheckPassword(password string) bool {
	return password == u.password
}

// Setters for private fields
func (u *User) UpdatePassword(newPassword string) {
	u.password = newPassword
}

func abstraction() {
	fmt.Println("--------------------------------------------------------Abstraction--------------------------------------------------------")
	user := NewUser("Alice", "alice@abc.com", "password123")
	var auth Authenticator = user

	// Reading Public Fields directly
	fmt.Println("Reading Public Fields directly without any getters")
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)

	// Reading Private Fields using getters
	fmt.Println("\nReading Private Fields with getters")
	fmt.Printf("Login successful: %v\n", auth.CheckPassword("wrong"))

	fmt.Printf("\nUser Details: %v\n", user) // Here private fields will also be shown, but that doesn't mean encapsulation is broken.

	fmt.Println()
	fmt.Println()

	// Attempting to modify public and private fields
	fmt.Println("Updating Public Fields directly without any setters")
	user.Email = "alice_newid@abc.com"
	fmt.Printf("Email: %s\n", user.Email)

	fmt.Println("\nUpdating Private Fields using setter")
	fmt.Printf("Updating Password\n")
	auth.UpdatePassword("newpassword123")
	fmt.Printf("Login successful: %v\n", auth.CheckPassword("newpassword123"))

	fmt.Printf("\nUser Details: %v\n", user)

	fmt.Println()
	fmt.Println("Why is this called abstraction? Because it added an extra layer of interface that hides the password management and allows controlled access to sensitive data through methods, while still allowing direct access to public fields")
	fmt.Println("This is also encapsulation, because the password is private and only accessible through getters and setters")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")

}
