package solid

// ********************* Dependency Injection Principle *********************

import "fmt"

// DI Definition
// High-level modules should not depend on low-level modules. Both should depend on abstractions
type EmailSender_DI struct{}

func (e EmailSender_DI) Send(i Invoice) {
	fmt.Printf("Sending email to %s\n", i.Customer)
}

// BAD EXAMPLE of Dependency Injection (DI) //
// Here InvoiceService depends directly on a concrete type: EmailSender_DI. This is tightly coupled
// Suppose instead of sending emails, we want to send then text messages., then we need to modify the InvoiceService struct
type B_InvoiceService struct {
	emailSender EmailSender_DI
}

func (s B_InvoiceService) SendInvoice(i Invoice) {
	// Tight coupling
	s.emailSender.Send(i)
}

// GOOD EXAMPLE of Dependency Injection (DI) //
// G_InvoiceService uses an interface SendMessage, which allows it to depend on any type that implements SendMessage.
// This makes it flexible and open for extension without modifying the existing code.
type SendMessage interface {
	Send(i Invoice)
}

type G_InvoiceService struct {
	mailer SendMessage
}

func (s G_InvoiceService) SendInvoice(i Invoice) {
	s.mailer.Send(i)
}
