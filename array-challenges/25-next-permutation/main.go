package main

import "fmt"

func nextPermutation(arr []int) []int {
	n := len(arr)
	i := n - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	reverseFrom(arr, i+1)
	return arr
}

func reverseFrom(arr []int, start int) {
	for i, j := start, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("Next permutation:", nextPermutation(arr))
}
