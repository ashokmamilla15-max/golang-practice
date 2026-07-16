package main

import (
	"fmt"
	"sort"
)

func kthLargest(arr []int, k int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	return sorted[k-1]
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println("3rd largest element:", kthLargest(arr, 3))
}
