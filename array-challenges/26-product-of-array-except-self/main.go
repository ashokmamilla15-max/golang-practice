package main

import "fmt"

func productExceptSelf(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * arr[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= arr[i]
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("Product except self:", productExceptSelf(arr))
}
