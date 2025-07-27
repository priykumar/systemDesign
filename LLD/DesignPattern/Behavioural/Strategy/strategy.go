package strategy

import "fmt"

// Strategy Pattern Example
// This pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable.
// In this example, we create a payment processing system with different payment strategies.
// The PaymentStrategy interface defines the method for payment processing.
// The UpiPayment and CreditCardPayment structs implement the PaymentStrategy interface.
// The PaymentProcessor struct uses a PaymentStrategy to process payments.
// The client can choose which payment strategy to use at runtime.

type PaymentStrategy interface {
	Pay(amount float64) string
}

// Payment strategy 1
type UpiPayment struct{}

func (u *UpiPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using UPI", amount)
}

// Payment strategy 2
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using Credit Card", amount)
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (p *PaymentProcessor) Checkout(amount float64) string {
	return p.strategy.Pay(amount)
}

func Strategy_DP() {
	// Use UPI
	strategy1 := &UpiPayment{}
	upi := NewPaymentProcessor(strategy1)
	fmt.Println(upi.Checkout(500))

	// Use Credit Card
	strategy2 := &CreditCardPayment{}
	card := NewPaymentProcessor(strategy2)
	fmt.Println(card.Checkout(1500))
}
