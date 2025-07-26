package factory

import "fmt"

// Factory Design Pattern
// In this example, we create a simple factory that produces different types of database connections based on the input type.

type Database interface {
	Connect() string
}

type MySQL struct{}

func (m MySQL) Connect() string { return "Connected to MySQL" }

type PostgreSQL struct{}

func (p PostgreSQL) Connect() string { return "Connected to PostgreSQL" }

// Factory
func NewDatabase(dbType string) Database {
	switch dbType {
	case "mysql":
		return &MySQL{}
	case "postgresql":
		return &PostgreSQL{}
	default:
		return &MySQL{} // default
	}
}

// Usage
func Factory_DP() {
	db := NewDatabase("postgresql")
	fmt.Println(db.Connect())
}
