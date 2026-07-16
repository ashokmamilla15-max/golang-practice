package main

import "fmt"

func pivotIndex(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	leftSum := 0
	for i, v := range arr {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}

func main() {
	arr := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("Pivot index:", pivotIndex(arr))
}
