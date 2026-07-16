// A type switch compares the concrete type stored in an interface value
// against multiple types in a single switch, using the form
// `switch v := i.(type) { case T: ... }`.
package main

import "fmt"

func describe(i interface{}) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	case []int:
		return fmt.Sprintf("slice of int with %d elements", len(v))
	default:
		return fmt.Sprintf("unhandled type %T", v)
	}
}

func main() {
	fmt.Println(describe(42))
	fmt.Println(describe("go"))
	fmt.Println(describe(true))
	fmt.Println(describe([]int{1, 2, 3}))
	fmt.Println(describe(3.14))
}
