// Methods are functions with a receiver argument, associating them with a
// type. A pointer receiver can modify the original value and avoids
// copying; a value receiver operates on a copy.
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("area:", rect.Area())

	rect.Scale(2)
	fmt.Println("after scale:", rect, "new area:", rect.Area())
}
