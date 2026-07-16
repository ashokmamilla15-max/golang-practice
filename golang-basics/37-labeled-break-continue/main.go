// Labels give `break` and `continue` a target outer loop, useful for
// escaping or skipping nested loops that a plain break/continue (which
// only affects the innermost loop) cannot reach.
package main

import "fmt"

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer
			}
			if i == 2 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
}
