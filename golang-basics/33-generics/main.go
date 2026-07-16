// Generics (added in Go 1.18) let functions and types operate over any
// type that satisfies a constraint, avoiding duplicated code or
// interface{}-based workarounds. `any` is an alias for interface{}.
package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5}))

	doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	fmt.Println(doubled)
}
