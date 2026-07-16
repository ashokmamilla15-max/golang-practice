package main

import "fmt"

func maxSubArraySum(arr []int) int {
	maxSum := arr[0]
	currentSum := arr[0]
	for _, v := range arr[1:] {
		currentSum = max(v, currentSum+v)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum subarray sum:", maxSubArraySum(arr))
}
