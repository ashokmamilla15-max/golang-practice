// Variables in Go can be declared with the `var` keyword (with or without
// an explicit type) or with the short declaration operator `:=` inside
// functions. Uninitialized variables get the "zero value" for their type:
// 0 for numbers, "" for strings, false for bools, nil for pointers/slices/
// maps/interfaces/channels/functions.
package main

import "fmt"

func main() {
	var name string = "Gopher"
	var age = 5
	var isActive bool
	var score float64

	city := "Bengaluru"

	var x, y int = 1, 2
	var a, b = "hello", 42

	fmt.Println(name, age, isActive, score)
	fmt.Println(city)
	fmt.Println(x, y)
	fmt.Println(a, b)
}
