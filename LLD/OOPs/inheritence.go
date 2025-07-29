package main

import "fmt"

// This looks like inheritance, but it is actually composition
// Here Rect is a parent struct and Square is a child struct and Rect has features like Area() and Perimeter() which are inherited by Square
// Here, Square is able to access the methods of Rect(s.Area()) and also able to access its fields(s.l), hence it looks like inheritance
// It feels like inheritance, but it's actually composition with method promotion.

// This line:
// type Square struct {
// 	Rect
// }

// is syntactic sugar for:

// type Square struct {
// 	Rect Rect
// }

// BUT with an important twist:
// When you embed a struct (instead of naming the field), Go automatically promotes the fields and methods of the embedded struct to the outer struct.
// So you can access s.l, s.Area(), etc., as if they belong to Square.
// But under the hood, Go is doing delegation, not inheritance.

type S interface {
	Area() int
	Perimeter() int
}

type Rect struct {
	l int
	w int
}

type Square struct {
	Rect
}

func (r *Rect) Area() int {
	return r.l * r.w
}
func (r *Rect) Perimeter() int {
	return 2 * (r.l + r.w)
}

func not_inheritence() {
	s := &Square{
		Rect: Rect{
			l: 5,
			w: 10,
		},
	}
	fmt.Println(s.l)
	s.Area()
}
