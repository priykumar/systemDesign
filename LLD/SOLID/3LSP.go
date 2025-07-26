package solid

// ********************* Liskov Substitution Principle *********************

// LSP Definition
// Subtypes must be substitutable for their base types.
// Anywhere you use a parent type (like an interface), you should be able to replace it with any of its child types without breaking the behavior.
type TaxCalculator_LSP interface {
	Calculate(i Invoice) float64
}

type RegularTax struct{}
type CorporateTax struct{}

func (r RegularTax) Calculate(i Invoice) float64 {
	return i.Amount * 0.18
}

func (c CorporateTax) Calculate(i Invoice) float64 {
	return 0.15 * i.Amount
}

// BAD EXAMPLE of Liskov Substitution Principle (LSP) //
// Here TaxCalculator_LSP is a parent and RegularTax, CorporateTax, and InvalidTax implements TaxCalculator_LSP
// that means these are chidren of TaxCalculator_LSP.
// According to LSP, we should be able to replace parent(TaxCalculator_LSP) with any of its children without breaking the code, but
// InvalidTax panics and breaks the code.
// This violates the Liskov Substitution Principle because InvalidTax cannot be used interchangeably with
// with TaxCalculator_LSP without causing a runtime error.

// Violation
type InvalidTax struct{}

func (i InvalidTax) Calculate(inv Invoice) float64 {
	panic("should not be called") // breaks substitution!
}

// GOOD EXAMPLE of Liskov Substitution Principle (LSP) //
func (i InvalidTax) Calculate(inv Invoice) float64 {
	// Safe fallback behavior
	return 0.0
}
