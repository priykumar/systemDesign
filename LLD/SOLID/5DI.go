package solid

// ********************* Dependency Injection Principle *********************

import "fmt"

// DI Definition
// High-level modules should not depend on low-level modules. Both should depend on abstractions (interface)
type MySQLRepository struct{}

func (m *MySQLRepository) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

type UserService struct {
    repo *MySQLRepository
}

func (u *UserService) CreateUser(name string) {
    u.repo.Save(name) // tightly coupled to MySQL
}


// GOOD EXAMPLE of Dependency Injection (DI) //
// UserService is dependent on UserRepository(Abstraction/interface) rather than being dependent on MySQLRepository struct
// This makes it flexible and open for extension without modifying the existing code.
// Abstraction
type UserRepository interface {
    Save(data string)
}

// Low-level module
type MySQLRepository struct{}

func (m *MySQLRepository) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

// High-level module
type UserService struct {
    repo UserRepository // depends on abstraction
}

func (u *UserService) CreateUser(name string) {
    u.repo.Save(name)
}

// Main wiring
func main() {
    mysqlRepo := &MySQLRepository{}
    service := &UserService{repo: mysqlRepo}
    service.CreateUser("Priyanshu")
}

