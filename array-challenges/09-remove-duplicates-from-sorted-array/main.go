package main

import "fmt"

func removeDuplicatesSorted(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Println("Array without duplicates:", removeDuplicatesSorted(arr))
}
