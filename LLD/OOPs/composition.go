package main

import "fmt"

// Composition is a design principle where one struct contains another struct, allowing for code reuse and modularity.
// It is often referred to as "has-a" relationship, as opposed to inheritance which is "is-a" relationship.
// In this example, we will create a User struct and an Admin struct that "inherits" from User using composition.

// type User struct {
// 	Name     string // Public fields
// 	Email    string // Public fields
// 	password string // Private fields
// }

// "Inheriting" struct
type Admin struct {
	User            // Embedded struct (composition = inheritance)
	AccessLevel int // Extra field for Admin
}

// Admin-specific method
func (a Admin) ShowAdminInfo() {
	fmt.Printf("Admin name: %s, Email: %s, Level: %d\n", a.Name, a.Email, a.AccessLevel)
}

func composition() {
	fmt.Println("--------------------------------------------------------Inheritance--------------------------------------------------------")

	u := NewUser("Alice", "alice@admin.com")
	admin := Admin{
		User:        *u,
		AccessLevel: 5,
	}

	// Inherited public fields
	fmt.Println("Accessing public fields directly")
	fmt.Println("\tAdmin Name:", admin.Name)   // From User
	fmt.Println("\tAdmin Email:", admin.Email) // From User

	fmt.Println()

	// Inherited method
	fmt.Println("Accessing private fields through inherited method")
	fmt.Println("\tPassword correct?", admin.CheckPassword("admin123"))
	admin.UpdatePassword("defaultPassword", "newpass")
	fmt.Println("\tNew Password Valid?", admin.CheckPassword("newpass"))

	// ✅ Admin-specific method
	fmt.Println()
	admin.ShowAdminInfo()

	fmt.Println()
	fmt.Println("Why is this called inheritance? Because Admin struct inherits fields and methods from User struct")

	fmt.Println("-------------------------------------------------------------------------------")
}
