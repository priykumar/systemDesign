package solid

// ********************* Interface Segregation Principle *********************

import "fmt"

// ISP Definition
// A client should not be forced to depend on interfaces it does not use.
// This means that interfaces should be small and specific to the client’s needs.

type SimplePrinter struct{}

// BAD EXAMPLE of Interface Segregation Principle (ISP) //
// Here SimplePrinter is forced to implement methods it does not need.
// It implements all methods of InvoiceProcessor, but it only needs PrintInvoice.
// This violates the Interface Segregation Principle because SimplePrinter is forced to implement methods it does
type InvoiceProcessor interface {
	CalculateTax(i Invoice) float64
	PrintInvoice(i Invoice)
	SendEmail(i Invoice)
	GeneratePDF(i Invoice)
}

func (s SimplePrinter) CalculateTax(i Invoice) float64 {
	panic("not needed")
}
func (s SimplePrinter) PrintInvoice(i Invoice) {
	fmt.Println("Printing...")
}
func (s SimplePrinter) SendEmail(i Invoice) {
	panic("not needed")
}
func (s SimplePrinter) GeneratePDF(i Invoice) {
	panic("not needed")
}

// GOOD EXAMPLE of Interface Segregation Principle (ISP) //
// Here we define smaller, more specific interfaces.
// Each interface has only the methods that are relevant to the client, so SimplePrinter only implements PrintInvoice_ISP.
// This allows clients to depend only on the methods they actually use, following the Interface Segregation Principle.
type TaxCalculator_ISP interface {
	CalculateTax(i Invoice) float64
}

type Printer interface {
	PrintInvoice_ISP(i Invoice)
}

type EmailSender_ISP interface {
	SendEmail(i Invoice)
}

type PDFGenerator interface {
	GeneratePDF(i Invoice)
}

func (s SimplePrinter) PrintInvoice_ISP(i Invoice) {
	fmt.Println("Printing simple invoice...")
}
