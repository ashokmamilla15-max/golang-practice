// A pointer holds the memory address of a value. `&x` yields a pointer to
// x, and `*p` dereferences a pointer to access or modify the value it
// points to. Passing a pointer to a function lets the function mutate the
// caller's original value.
package main

import "fmt"

func increment(n *int) {
	*n++
}

func main() {
	x := 10
	p := &x

	fmt.Println("value:", *p)
	*p = 20
	fmt.Println("after *p = 20:", x)

	increment(&x)
	fmt.Println("after increment:", x)

	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr == nil)
}
