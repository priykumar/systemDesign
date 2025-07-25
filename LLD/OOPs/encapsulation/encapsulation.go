package encapsulation

// Encapsulation is defined as the wrapping up of data under a single unit
// Encapsulation is exactly like having a simple switch that hides all the complex electrical work happening behind the scenes

type User struct {
	Name     string // Public fields - direct access allowed
	Email    string // Public fields - direct access allowed
	password string // Private fields - controlled access only
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

// Getters for private fields
func CheckPassword(u *User, password string) bool {
	return password == u.password
}

// Setters for private fields
func (u *User) UpdatePassword(newPassword string) {
	u.password = newPassword
}

// Setters for private fields
func UpdatePassword(u *User, newPassword string) {
	u.password = newPassword
}
