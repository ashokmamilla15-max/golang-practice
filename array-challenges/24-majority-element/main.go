package main

import "fmt"

func majorityElement(arr []int) int {
	count, candidate := 0, 0
	for _, v := range arr {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority element:", majorityElement(arr))
}
