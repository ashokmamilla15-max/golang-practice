// Function literals (anonymous functions) can be assigned to variables,
// passed as arguments, or invoked immediately (IIFE).
package main

import "fmt"

func main() {
	square := func(n int) int {
		return n * n
	}
	fmt.Println(square(5))

	func() {
		fmt.Println("invoked immediately")
	}()

	apply := func(n int, f func(int) int) int {
		return f(n)
	}
	fmt.Println(apply(4, func(n int) int { return n * n * n }))
}
