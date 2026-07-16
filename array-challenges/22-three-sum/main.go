package main

import (
	"fmt"
	"sort"
)

func threeSum(arr []int) [][]int {
	sort.Ints(arr)
	result := make([][]int, 0)
	n := len(arr)
	for i := 0; i < n-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			switch {
			case sum == 0:
				result = append(result, []int{arr[i], arr[left], arr[right]})
				left++
				right--
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			case sum < 0:
				left++
			default:
				right--
			}
		}
	}
	return result
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Triplets that sum to zero:", threeSum(arr))
}
