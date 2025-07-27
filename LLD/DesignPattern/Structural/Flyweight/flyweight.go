package flyweight

import "fmt"

// Flyweight Pattern Example
// This pattern is used to minimize memory usage by sharing common parts of objects.
// In this example, we create a character factory that manages characters with different fonts.
// The Character struct represents the intrinsic state (shared part) of a character.
// The CharacterFactory manages the creation and reuse of characters.
// The Character struct has intrinsic properties (symbol, font) and methods to draw the character.
// The CharacterFactory checks if a character with the same symbol and font already exists.
// If it exists, it returns the existing character; otherwise, it creates a new one.
// The Flyweight pattern is useful when dealing with a large number of similar objects that share some common state.

// Flyweight (shared part)
type Character struct {
	symbol string // intrinsic
	font   string
}

func (c *Character) Draw(x, y int) {
	fmt.Printf("Drawing '%s' in font '%s' at (%d, %d)\n", c.symbol, c.font, x, y)
}

type CharacterFactory struct {
	characters map[string]*Character
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{characters: make(map[string]*Character)}
}

func (f *CharacterFactory) GetCharacter(symbol, font string) *Character {
	key := symbol + font
	if ch, exists := f.characters[key]; exists {
		return ch
	}
	ch := &Character{symbol, font}
	f.characters[key] = ch
	return ch
}

func Flyweight_DP() {
	factory := NewCharacterFactory()

	// Create 'A' in Arial font
	a1 := factory.GetCharacter("A", "Arial")
	a1.Draw(1, 2)

	// Reuse 'A' in Arial font
	a2 := factory.GetCharacter("A", "Arial")
	a2.Draw(3, 4)

	// Create 'B' with Arial
	b := factory.GetCharacter("B", "Arial")
	b.Draw(5, 6)

	// Create 'A' in Times font
	a3 := factory.GetCharacter("A", "Times")
	a3.Draw(7, 8)
}
