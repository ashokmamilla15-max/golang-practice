// A parameter prefixed with `...T` accepts a variable number of arguments,
// available inside the function as a []T slice. An existing slice can be
// passed with the `slice...` spread syntax.
package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))

	nums := []int{4, 5, 6}
	fmt.Println(sum(nums...))
}
