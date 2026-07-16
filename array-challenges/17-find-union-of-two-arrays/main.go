package main

import "fmt"

func union(a, b []int) []int {
	set := make(map[int]bool)
	result := make([]int, 0)
	for _, v := range a {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}
	for _, v := range b {
		if !set[v] {
			set[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println("Union:", union(a, b))
}
