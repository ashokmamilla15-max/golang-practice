package main

import "fmt"

func mergeSortedArrays(a, b []int) []int {
	merged := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}
	merged = append(merged, a[i:]...)
	merged = append(merged, b[j:]...)
	return merged
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	fmt.Println("Merged array:", mergeSortedArrays(a, b))
}
