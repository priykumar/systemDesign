package main

import "fmt"

// Polymorphism is the ability to present the same interface for different data types.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type circle struct {
	radius float64
}
type rectangle struct {
	length float64
	width  float64
}
type square struct {
	side float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}
func (r rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (s square) Area() float64 {
	return s.side * s.side
}

func polymorphism() {
	fmt.Println("--------------------------------------------------------Polymorphism--------------------------------------------------------")
	var s Shape
	c := circle{radius: 5}
	fmt.Println("Using Circle struct")
	fmt.Printf("\tShape: %v Type: %T\n", c, c)
	fmt.Printf("\tCircle Area: %.2f\n", c.Area())
	fmt.Printf("\tCircle Perimeter: %.2f\n", c.Perimeter())

	fmt.Println()

	r := rectangle{length: 4, width: 6}
	fmt.Println("Using Rectangle struct")
	fmt.Printf("\tShape: %v Type: %T\n", r, r)
	fmt.Printf("\tRectangle Area: %.2f\n", r.Area())
	fmt.Printf("\tRectangle Perimeter: %.2f\n", r.Perimeter())

	fmt.Println()

	square := square{side: 3}
	fmt.Println("Using Square struct")
	fmt.Printf("\tShape: %v Type: %T\n", square, square)
	fmt.Printf("\tSquare Area: %.2f\n", square.Area())

	fmt.Println()

	s = c
	fmt.Println("Using Shape interface, as Cirle implements Shape interface because Circle has defines all the methods of Shape interface")
	fmt.Printf("\tShape: %v Type: %T\n", s, s)
	fmt.Printf("\tCircle Area: %.2f\n", s.Area())
	fmt.Printf("\tCircle Perimeter: %.2f\n", s.Perimeter())

	fmt.Println()

	s = r
	fmt.Println("Using Shape interface, as Rectangle implements Shape interface because Rectangle has defines all the methods of Shape interface")
	fmt.Printf("\tShape: %v Type: %T\n", s, s)
	fmt.Printf("\tRectangle Area: %.2f\n", s.Area())
	fmt.Printf("\tRectangle Perimeter: %.2f\n", s.Perimeter())

	fmt.Println()
	fmt.Println("Cant use square as Shape interface, because Square does not implement Perimeter method of Shape interface")
	// s = square

	fmt.Println()
	fmt.Println("Why is this called polymorphism? Because we can use the same interface (Shape) to refer to different types of shapes (Circle, Rectangle) and call their methods without knowing the exact type of shape at compile time")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------")
}
