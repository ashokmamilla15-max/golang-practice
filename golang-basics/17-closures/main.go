// A closure is a function value that references variables from outside its
// body. Each closure captures its own reference to those variables, so
// separate closures created from the same generator maintain independent
// state.
package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c1 := counter()
	c2 := counter()

	fmt.Println(c1(), c1(), c1())
	fmt.Println(c2())
}
