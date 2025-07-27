package observer

import (
	"testing"
)

// --- Mock Customer ---
type MockCustomer struct {
	Name     string
	Notified []string
}

func (m *MockCustomer) Update(productName string) {
	m.Notified = append(m.Notified, productName)
}

func (m *MockCustomer) GetName() string {
	return m.Name
}

func TestObserverPattern(t *testing.T) {
	// Create product
	product := NewProduct("iPhone 15")

	// Create customers
	alice := &MockCustomer{Name: "Alice"}
	bob := &MockCustomer{Name: "Bob"}

	// Register customers
	product.Register(alice)
	product.Register(bob)

	// Make product in stock (should notify both)
	product.SetInStock(true)

	// Assert both received notification
	if len(alice.Notified) != 1 || alice.Notified[0] != "iPhone 15" {
		t.Errorf("Alice was not notified correctly: %v", alice.Notified)
	}
	if len(bob.Notified) != 1 || bob.Notified[0] != "iPhone 15" {
		t.Errorf("Bob was not notified correctly: %v", bob.Notified)
	}

	// Deregister Bob
	product.Deregister(bob)

	// Restock again
	product.SetInStock(true)

	// Alice should be notified again
	if len(alice.Notified) != 2 {
		t.Errorf("Alice should have 2 notifications, got %d", len(alice.Notified))
	}

	// Bob should still have only 1
	if len(bob.Notified) != 1 {
		t.Errorf("Bob should have only 1 notification, got %d", len(bob.Notified))
	}
}
