// A struct groups named fields into a single type. Structs are value
// types: assigning or passing one copies all its fields.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{"Bob", 25} // positional fields, order must match

	fmt.Println(p1, p2)

	p3 := p1
	p3.Age = 99
	fmt.Println("original unaffected:", p1)
	fmt.Println("copy modified:", p3)

	p4 := &Person{Name: "Carol", Age: 28}
	p4.Age = 29 // Go auto-dereferences: shorthand for (*p4).Age = 29
	fmt.Println(*p4)
}
