package main

import "fmt"

func findMissingNumber(arr []int, n int) int {
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, v := range arr {
		actualSum += v
	}
	return expectedSum - actualSum
}

func main() {
	arr := []int{1, 2, 4, 5, 6}
	fmt.Println("Missing number:", findMissingNumber(arr, 6))
}
