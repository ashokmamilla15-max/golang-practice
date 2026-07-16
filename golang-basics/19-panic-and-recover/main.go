// `panic` stops normal execution and begins unwinding the stack, running
// deferred calls along the way. `recover`, called inside a deferred
// function, regains control of a panicking goroutine and returns the
// panic value, letting a program handle the error instead of crashing.
package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = a / b
	return
}

func main() {
	result, err := safeDivide(10, 0)
	fmt.Println(result, err)

	result, err = safeDivide(10, 2)
	fmt.Println(result, err)
}
