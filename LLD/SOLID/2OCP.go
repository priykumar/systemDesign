package solid

// ********************* Open Close Principle *********************

import "fmt"

// OCP Definition
// Software entities should be open for extension, but closed for modification

func CalculateTax(i Invoice) float64 {
	return i.Amount * 0.18
}

// BAD EXAMPLE of Open/Closed Principle (OCP) //
// Here Invoice struct is tightly coupled with the way it prints invoices.
// You must modify PrintInvoice() every time you want to support a new format.
// This violates the Open/Closed Principle because the function is not open for extension without modification
func B_PrintInvoice(i Invoice, format string) {
	switch format {
	case "text":
		fmt.Printf("Invoice: %s - ₹%.2f, Tax: ₹%.2f\n", i.Customer, i.Amount, CalculateTax(i))
	case "pdf":
		fmt.Println("Generating PDF invoice...") // imagine PDF logic here
	case "html":
		fmt.Println("Rendering HTML invoice...") // imagine HTML logic here
	default:
		fmt.Println("Unsupported format")
	}
}

// GOOD EXAMPLE of Open/Closed Principle (OCP) //
// Here we define an interface for printing invoices. Each format has its own implementation.
// This allows us to add new formats without modifying existing code.
// You don’t touch existing code to add new formats. Just create a new struct implementing Formatter.

// OCP-friendly: Formatter interface
type Formatter interface {
	Print(i Invoice)
}

// Text format
type TextFormatter struct{}

func (t TextFormatter) Print(i Invoice) {
	fmt.Printf("Invoice (Text): %s - ₹%.2f, Tax: ₹%.2f\n", i.Customer, i.Amount, CalculateTax(i))
}

// PDF format
type PDFFormatter struct{}

func (p PDFFormatter) Print(i Invoice) {
	fmt.Println("Generating PDF invoice...")
}

// HTML format
type HTMLFormatter struct{}

func (h HTMLFormatter) Print(i Invoice) {
	fmt.Println("Rendering HTML invoice...")
}

// Unified print function
func G_PrintInvoice(i Invoice, formatter Formatter) {
	formatter.Print(i)
}
