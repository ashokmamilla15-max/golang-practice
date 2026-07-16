// Slices are dynamically-sized views over an underlying array. They have a
// length and a capacity, grow with append, and are reference types: two
// slices can share the same backing array.
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))

	s = append(s, 4, 5)
	fmt.Println("after append:", s)

	sub := s[1:3]
	fmt.Println("sub-slice:", sub)

	made := make([]int, 3, 10)
	fmt.Println("make(len=3, cap=10):", made, len(made), cap(made))

	a := []int{1, 2, 3}
	b := a[:]
	b[0] = 100
	fmt.Println("shared backing array:", a, b)

	dst := make([]int, len(a))
	copy(dst, a)
	dst[0] = -1
	fmt.Println("independent copy:", a, dst)
}
