package observer

import "fmt"

// Observer Pattern Example
// This pattern defines a one-to-many dependency between objects so that when one object changes state,
// all its dependents are notified and updated automatically.

// In this example, we create a product availability notification system.
// The Product interface defines methods for registering, deregistering, and notifying customers.
// The Customer interface defines methods for receiving updates.
// The FlipkartProduct struct implements the Product interface and manages a list of customers.
// The FlipkartCustomer struct implements the Customer interface and receives notifications when a product is back in stock.
// The Observer pattern is useful for implementing event-driven systems where changes in one part of the system
type Customer interface {
	Update(productName string)
	GetName() string
}

type Product interface {
	Register(customer Customer)
	Deregister(customer Customer)
	NotifyAll()
	SetInStock(status bool)
}

// FlipkartCustomer implements the Customer interface
// It represents a customer who wants to be notified when a product is back in stock.
type FlipkartCustomer struct {
	Name string
}

func (c *FlipkartCustomer) Update(productName string) {
	fmt.Printf("Hello %s! '%s' is now back in stock!\n", c.Name, productName)
}

func (c *FlipkartCustomer) GetName() string {
	return c.Name
}

// FlipkartProduct implements the Product interface
// It represents a product on Flipkart that customers can subscribe to for stock updates.
type FlipkartProduct struct {
	Name        string
	inStock     bool
	subscribers []Customer
}

func NewProduct(name string) *FlipkartProduct {
	return &FlipkartProduct{Name: name}
}

func (p *FlipkartProduct) Register(c Customer) {
	p.subscribers = append(p.subscribers, c)
}

func (p *FlipkartProduct) Deregister(c Customer) {
	for i, customer := range p.subscribers {
		if customer.GetName() == c.GetName() {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
}

func (p *FlipkartProduct) NotifyAll() {
	for _, customer := range p.subscribers {
		customer.Update(p.Name)
	}
}

func (p *FlipkartProduct) SetInStock(status bool) {
	p.inStock = status
	if status {
		fmt.Printf("'%s' is now in stock.\n", p.Name)
		p.NotifyAll()
	}
}

func Observer_DP() {
	product := NewProduct("iPhone 15")

	// Customers
	alice := &FlipkartCustomer{Name: "Alice"}
	bob := &FlipkartCustomer{Name: "Bob"}

	// Customers subscribe
	product.Register(alice)
	product.Register(bob)

	// Product becomes available
	product.SetInStock(true)

	// Bob unsubscribes
	product.Deregister(bob)

	// Restocked again
	product.SetInStock(true)
}
