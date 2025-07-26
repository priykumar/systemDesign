package solid

// ********************* Single Responsibility Principle *********************

import "fmt"

// SRP Definition
// A struct or function should have only one reason to change — it should do only one job.
type Invoice struct {
	Customer string
	Amount   float64
}

// BAD EXAMPLE of Single Responsibility Principle (SRP) //
// Here Invoice struct is handling multiple responsibilities. It's involved in
// calculating tax, printing the invoice, and sending an email
func (i Invoice) CalculateTax() float64 {
	return i.Amount * 0.18
}

func (i Invoice) PrintInvoice() {
	fmt.Printf("Invoice for %s\nAmount: %.2f\nTax: %.2f\n", i.Customer, i.Amount, i.CalculateTax())
}

func (i Invoice) SendEmail() {
	fmt.Printf("Sending email to %s...\n", i.Customer)
}

// GOOD EXAMPLE of Single Responsibility Principle (SRP) //
// Here Invoice struct is focused on representing the invoice data, while
// TaxCalculator, InvoicePrinter, and EmailSender are separate types that handle
// specific responsibilities: calculating tax, printing the invoice, and sending emails.
// This adheres to the Single Responsibility Principle, making the code more maintainable and testable
type TaxCalculator struct{}

func (t TaxCalculator) Calculate(invoice Invoice) float64 {
	return invoice.Amount * 0.18
}

type InvoicePrinter struct{}

func (p InvoicePrinter) Print(invoice Invoice, tax float64) {
	fmt.Printf("Invoice for %s\nAmount: %.2f\nTax: %.2f\n", invoice.Customer, invoice.Amount, tax)
}

type EmailSender struct{}

func (e EmailSender) Send(invoice Invoice) {
	fmt.Printf("Sending email to %s...\n", invoice.Customer)
}
