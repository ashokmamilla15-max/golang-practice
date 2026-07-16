// Custom error types implement the error interface, letting callers carry
// structured information. `errors.As` (paired with wrapping via
// fmt.Errorf's %w verb) lets callers inspect wrapped error chains.
package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

func findUser(id int) error {
	if id != 1 {
		return &NotFoundError{Resource: fmt.Sprintf("user %d", id)}
	}
	return nil
}

func loadProfile(id int) error {
	if err := findUser(id); err != nil {
		return fmt.Errorf("loadProfile: %w", err)
	}
	return nil
}

func main() {
	err := loadProfile(2)
	fmt.Println(err)

	var nfErr *NotFoundError
	if errors.As(err, &nfErr) {
		fmt.Println("resource missing:", nfErr.Resource)
	}
}
