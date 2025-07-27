package flyweight

import (
	"testing"
)

func TestCharacterFactory_ReusesCharacter(t *testing.T) {
	factory := NewCharacterFactory()

	// Create first 'A' with Arial
	char1 := factory.GetCharacter("A", "Arial")
	// Create second 'A' with Arial — should reuse the same instance
	char2 := factory.GetCharacter("A", "Arial")

	if char1 != char2 {
		t.Errorf("Expected reused instance for 'A' with Arial, but got different instances")
	}
}

func TestCharacterFactory_CreatesNewCharacterForDifferentFont(t *testing.T) {
	factory := NewCharacterFactory()

	charArial := factory.GetCharacter("A", "Arial")
	charTimes := factory.GetCharacter("A", "Times")

	if charArial == charTimes {
		t.Errorf("Expected different instances for same symbol with different fonts")
	}
}

func TestCharacter_Draw(t *testing.T) {
	char := &Character{symbol: "X", font: "Courier"}
	char.Draw(10, 20)
}
