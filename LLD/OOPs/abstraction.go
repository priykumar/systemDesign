package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unicode"
)

// Abstraction is the concept of hiding the complex implementation details
// Here, we are hiding the complexity of password management and validation behind an interface
type Authenticator interface {
	CheckPassword(password string) bool
	UpdatePassword(currPassword, newPassword string)
}

type User struct {
	Name     string // Public fields
	Email    string // Public fields
	password string // Private fields
}

func NewUser(name, email string) *User {
	return &User{
		Name:     name,
		Email:    email,
		password: "defaultPassword",
	}
}

// Getters for private fields
func (u *User) CheckPassword(password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	gen_hash := hex.EncodeToString(hash.Sum(nil))
	return gen_hash == u.password
}

// Complexity of password management and validation is hidden behind this interface, hence abstraction
func (u *User) UpdatePassword(currPassword, newPassword string) {
	// validate current password
	if u.CheckPassword(currPassword) {
		fmt.Println("Current password is incorrect")
		return
	}

	// validate new password strength
	if len(newPassword) < 8 {
		fmt.Println("New password must be at least 8 characters long")
		return
	}
	// check for at least one lowercase, one uppercase, one digit, and one special character
	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for _, ch := range newPassword {
		switch {
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !(hasLower && hasUpper && hasDigit && hasSpecial) {
		fmt.Println("New password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character")
		return
	}

	// generate a secure hash for the new password
	hash := sha256.New()
	hash.Write([]byte(newPassword))
	u.password = hex.EncodeToString(hash.Sum(nil))
}

func abstraction() {
	fmt.Println("--------------------------------------------------------Abstraction--------------------------------------------------------")
	user := NewUser("Alice", "alice@abc.com")
	var auth Authenticator = user

	fmt.Printf("\nUser Details: %v\n", user) // Here private fields will also be shown, but that doesn't mean encapsulation is broken.

	fmt.Println()

	fmt.Printf("Updating Password\n")
	auth.UpdatePassword("defaultPassword", "newPassword@123")
	fmt.Printf("Login successful: %v\n", auth.CheckPassword("newPassword@123"))
	auth.UpdatePassword("defaultPassword", "newnewPassword@123")
	fmt.Printf("Login successful: %v\n", auth.CheckPassword("newnewPassword@123"))

	fmt.Printf("\nUser Details: %v\n", user)

	fmt.Println()
	fmt.Println("Why is this called abstraction? Because it hides the complexity of password management and validation behind a simple interface.")
	fmt.Println("This is also encapsulation, because the password is private and only accessible through getters and setters")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")

}

// REAL WORLD EXAMPLE - TV Example
// Encapsulation is hiding the internals of TV like wiring, circuits ...etc
// Abstraction is hiding the complex functionalities like whats would happen in the backend you power on/off, volume up/down, channel change ...etc
