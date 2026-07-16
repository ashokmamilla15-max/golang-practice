package main

import "fmt"

func intersection(a, b []int) []int {
	set := make(map[int]bool)
	for _, v := range a {
		set[v] = true
	}
	result := make([]int, 0)
	seen := make(map[int]bool)
	for _, v := range b {
		if set[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 2, 4, 6}
	fmt.Println("Intersection:", intersection(a, b))
}
