package main

import "fmt"

func maxProductSubArray(arr []int) int {
	maxProd, minProd, result := arr[0], arr[0], arr[0]
	for _, v := range arr[1:] {
		if v < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(v, maxProd*v)
		minProd = min(v, minProd*v)
		result = max(result, maxProd)
	}
	return result
}

func main() {
	arr := []int{2, 3, -2, 4}
	fmt.Println("Maximum product subarray:", maxProductSubArray(arr))
}
