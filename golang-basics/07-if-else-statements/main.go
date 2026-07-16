// `if` statements can include an optional initialization statement scoped
// to the if/else chain, e.g. `if x := compute(); x > 0 { ... }`.
package main

import "fmt"

func classify(n int) string {
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}

func main() {
	fmt.Println(classify(-5), classify(0), classify(5))

	if result := 10 * 2; result > 15 {
		fmt.Println("result is large:", result)
	}
}
