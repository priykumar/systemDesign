package decorator

import "testing"

func TestPizzaWithTomatoOnly(t *testing.T) {
	pizza := &tomatoTopping{
		pizza: &Pizza{},
	}

	expectedPrice := 150 + 50
	expectedConstituents := "pizza bread + tomato"

	if pizza.getPrice() != expectedPrice {
		t.Errorf("TomatoOnly: Expected price: %d, got: %d", expectedPrice, pizza.getPrice())
	}
	if pizza.getConstituents() != expectedConstituents {
		t.Errorf("TomatoOnly: Expected constituents: %s, got: %s", expectedConstituents, pizza.getConstituents())
	}
}

func TestPizzaWithCheeseOnly(t *testing.T) {
	pizza := &cheeseTopping{
		pizza: &Pizza{},
	}

	expectedPrice := 150 + 70
	expectedConstituents := "pizza bread + cheese"

	if pizza.getPrice() != expectedPrice {
		t.Errorf("CheeseOnly: Expected price: %d, got: %d", expectedPrice, pizza.getPrice())
	}
	if pizza.getConstituents() != expectedConstituents {
		t.Errorf("CheeseOnly: Expected constituents: %s, got: %s", expectedConstituents, pizza.getConstituents())
	}
}

func TestPizzaWithCheeseAndTomato(t *testing.T) {
	pizza := &tomatoTopping{
		pizza: &cheeseTopping{
			pizza: &Pizza{},
		},
	}

	expectedPrice := 150 + 70 + 50
	expectedConstituents := "pizza bread + cheese + tomato"

	if pizza.getPrice() != expectedPrice {
		t.Errorf("Cheese+Tomato: Expected price: %d, got: %d", expectedPrice, pizza.getPrice())
	}
	if pizza.getConstituents() != expectedConstituents {
		t.Errorf("Cheese+Tomato: Expected constituents: %s, got: %s", expectedConstituents, pizza.getConstituents())
	}
}

func TestPizzaWithAllToppings(t *testing.T) {
	pizza := &chickenTopping{
		pizza: &tomatoTopping{
			pizza: &cheeseTopping{
				pizza: &Pizza{},
			},
		},
	}

	expectedPrice := 150 + 70 + 50 + 100
	expectedConstituents := "pizza bread + cheese + tomato + chicken"

	if pizza.getPrice() != expectedPrice {
		t.Errorf("AllToppings: Expected price: %d, got: %d", expectedPrice, pizza.getPrice())
	}
	if pizza.getConstituents() != expectedConstituents {
		t.Errorf("AllToppings: Expected constituents: %s, got: %s", expectedConstituents, pizza.getConstituents())
	}
}
