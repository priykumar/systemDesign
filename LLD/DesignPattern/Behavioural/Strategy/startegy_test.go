package strategy

import "testing"

func TestUpiPayment(t *testing.T) {
	upi := &UpiPayment{}
	result := upi.Pay(500)
	expected := "Paid ₹500.00 using UPI"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestCreditCardPayment(t *testing.T) {
	card := &CreditCardPayment{}
	result := card.Pay(1500)
	expected := "Paid ₹1500.00 using Credit Card"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestPaymentProcessorWithUpi(t *testing.T) {
	processor := NewPaymentProcessor(&UpiPayment{})
	result := processor.Checkout(300)
	expected := "Paid ₹300.00 using UPI"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestPaymentProcessorWithCreditCard(t *testing.T) {
	processor := NewPaymentProcessor(&CreditCardPayment{})
	result := processor.Checkout(800)
	expected := "Paid ₹800.00 using Credit Card"

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
