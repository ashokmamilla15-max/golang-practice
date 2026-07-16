package main

import "fmt"

func subarraySum(arr []int, k int) int {
	count := 0
	sum := 0
	prefixSums := map[int]int{0: 1}
	for _, v := range arr {
		sum += v
		count += prefixSums[sum-k]
		prefixSums[sum]++
	}
	return count
}

func main() {
	arr := []int{1, 1, 1}
	fmt.Println("Number of subarrays with sum 2:", subarraySum(arr, 2))
}
