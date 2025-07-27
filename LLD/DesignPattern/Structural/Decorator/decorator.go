package decorator

import "fmt"

// Decorator Pattern Example
// This pattern allows adding new functionality to an object dynamically without altering its structure.

// Here, we create a pizza with different toppings using decorators.
// The base pizza interface defines the methods to get the price and constituents of the pizza.
// Each topping decorator implements the same interface and adds its own functionality.
// pizzaInterface defines the methods for pizza
// The Pizza struct implements pizzaInterface and represents the base pizza.
// The topping decorators (tomatoTopping, cheeseTopping, chickenTopping) implement the same interface and wrap the base pizza to add their own price and constituents.

type pizzaInterface interface {
	getPrice() int
	getConstituents() string
}

// Base pizza struct
type Pizza struct {
}

func (v *Pizza) getPrice() int {
	return 150
}
func (v *Pizza) getConstituents() string {
	pizzaMaterial := "pizza bread"
	return pizzaMaterial
}

// Topping decorators
// Each topping decorator wraps the base pizza and adds its own price and constituents.

// Decorator for tomato topping
type tomatoTopping struct {
	pizza pizzaInterface
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 50
}
func (t *tomatoTopping) getConstituents() string {
	pizzaMaterial := t.pizza.getConstituents() + " + tomato"
	return pizzaMaterial
}

// Decorator for cheese topping
type cheeseTopping struct {
	pizza pizzaInterface
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 70
}
func (c *cheeseTopping) getConstituents() string {
	pizzaMaterial := c.pizza.getConstituents() + " + cheese"
	return pizzaMaterial
}

// Decorator for chicken topping
type chickenTopping struct {
	pizza pizzaInterface
}

func (c *chickenTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 100
}
func (c *chickenTopping) getConstituents() string {
	pizzaMaterial := c.pizza.getConstituents() + " + chicken"
	return pizzaMaterial
}

func Decorator_DP() {
	// Create a pizza with cheese and tomato toppings
	vegPizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: &cheeseTopping{
			pizza: &Pizza{},
		},
	}

	// Create a non-veg pizza with chicken and tomato toppings
	nonVegPizza := &chickenTopping{
		pizza: &tomatoTopping{
			pizza: &cheeseTopping{
				pizza: &Pizza{},
			},
		},
	}

	fmt.Println("Veg Pizza Price:", vegPizzaWithCheeseAndTomato.getPrice())
	fmt.Println("Veg Pizza Includes:", vegPizzaWithCheeseAndTomato.getConstituents())

	fmt.Println("Non-Veg Pizza Price:", nonVegPizza.getPrice())
	fmt.Println("Non-Veg Pizza Includes:", nonVegPizza.getConstituents())
}
