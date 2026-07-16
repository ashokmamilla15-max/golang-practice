package main

import "fmt"

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	arr := []int{4, 2, 4, 1, 2, 5, 1}
	fmt.Println("Array without duplicates:", removeDuplicates(arr))
}
