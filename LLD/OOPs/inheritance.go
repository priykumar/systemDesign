package main

import "fmt"

// "Inheriting" struct
type Admin struct {
	User            // Embedded struct (composition = inheritance)
	AccessLevel int // Extra field for Admin
}

// Admin-specific method
func (a Admin) ShowAdminInfo() {
	fmt.Printf("Admin name: %s, Email: %s, Level: %d\n", a.Name, a.Email, a.AccessLevel)
}

func inheritance() {
	fmt.Println("--------------------------------------------------------Inheritance--------------------------------------------------------")

	u := NewUser("Alice", "alice@admin.com", "admin123")
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
	admin.UpdatePassword("newpass")
	fmt.Println("\tNew Password Valid?", admin.CheckPassword("newpass"))

	// ✅ Admin-specific method
	fmt.Println()
	admin.ShowAdminInfo()

	fmt.Println()
	fmt.Println("Why is this called inheritance? Because Admin struct inherits fields and methods from User struct")

	fmt.Println("-------------------------------------------------------------------------------")
}
