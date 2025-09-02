package solid

// ********************* Liskov Substitution Principle *********************

// LSP Definition
// Subtypes must be substitutable for their base types.
// Anywhere you use a parent type (like an interface), you should be able to replace it with any of its child types without breaking the behavior.
// Base class
type Bird interface {
    Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
    fmt.Println("Sparrow flying...")
}

// Subclass violates LSP
type Ostrich struct{}

func (o Ostrich) Fly() {
    panic("Ostrich cannot fly!") // ❌ breaks expectation
}


// GOOD EXAMPLE of Liskov Substitution Principle (LSP) //
// Abstraction split
type Bird interface {
    Eat()
}

type FlyingBird interface {
    Bird
    Fly()
}

// Flying bird
type Sparrow struct{}

func (s Sparrow) Eat() {
    fmt.Println("Sparrow eating...")
}

func (s Sparrow) Fly() {
    fmt.Println("Sparrow flying...")
}

// Non-flying bird
type Ostrich struct{}

func (o Ostrich) Eat() {
    fmt.Println("Ostrich eating...")
}
