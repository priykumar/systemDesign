package encapsulation

// Encapsulation means hiding the internal data and allowing controlled access to it through methods(getter and setter).
// Its like a capsule that contains the data and methods to manipulate it, but hides the internal implementation details.
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
