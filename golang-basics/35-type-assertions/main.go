// A type assertion, x.(T), extracts the concrete value stored in an
// interface. The two-result form, v, ok := x.(T), reports success instead
// of panicking when the assertion fails.
package main

import "fmt"

func describe(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Println("string of length", len(s))
		return
	}
	if n, ok := i.(int); ok {
		fmt.Println("int doubled:", n*2)
		return
	}
	fmt.Println("unknown type")
}

func main() {
	describe("hello")
	describe(21)
	describe(3.14)
}
