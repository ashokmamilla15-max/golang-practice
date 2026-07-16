// Functions are declared with `func`, can take typed parameters, and
// return a value. Go supports named return values, which are pre-declared
// and returned automatically by a bare `return`.
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (quotient int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	quotient = a / b
	return
}

func main() {
	fmt.Println(add(2, 3))

	q, err := divide(10, 2)
	fmt.Println(q, err)

	q, err = divide(10, 0)
	fmt.Println(q, err)
}
