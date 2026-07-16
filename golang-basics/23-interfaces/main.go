// An interface defines a method set. Any type that implements all of an
// interface's methods satisfies it implicitly, with no explicit
// "implements" declaration needed.
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func printArea(s Shape) {
	fmt.Printf("%T area: %.2f\n", s, s.Area())
}

func main() {
	shapes := []Shape{Circle{Radius: 2}, Square{Side: 3}}
	for _, s := range shapes {
		printArea(s)
	}

	var empty interface{} = "anything can satisfy the empty interface"
	fmt.Println(empty)
}
