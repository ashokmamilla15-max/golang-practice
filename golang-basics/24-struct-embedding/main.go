// Go favors composition over inheritance. Embedding one struct (or
// interface) inside another promotes its fields and methods, so the outer
// type can use them as if they were its own.
package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

type Dog struct {
	Animal // embedded field
	Breed  string
}

func main() {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}

	fmt.Println(d.Name)        // promoted field
	fmt.Println(d.Speak())     // promoted method
	fmt.Println(d.Animal.Name) // still accessible explicitly
}
