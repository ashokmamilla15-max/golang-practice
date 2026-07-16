// Errors in Go are ordinary values that implement the `error` interface
// (a single Error() string method). Functions that can fail return an
// error as their last result, and callers check it explicitly instead of
// using exceptions.
package main

import (
	"errors"
	"fmt"
)

func parsePositive(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("value must not be negative")
	}
	return n, nil
}

func main() {
	if v, err := parsePositive(5); err == nil {
		fmt.Println("parsed:", v)
	}

	if _, err := parsePositive(-1); err != nil {
		fmt.Println("error:", err)
	}
}
