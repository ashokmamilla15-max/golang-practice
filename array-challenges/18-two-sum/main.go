package main

import "fmt"

func twoSum(arr []int, target int) (int, int) {
	seen := make(map[int]int)
	for i, v := range arr {
		if j, ok := seen[target-v]; ok {
			return j, i
		}
		seen[v] = i
	}
	return -1, -1
}

func main() {
	arr := []int{2, 7, 11, 15}
	i, j := twoSum(arr, 9)
	fmt.Println("Indices:", i, j)
}
