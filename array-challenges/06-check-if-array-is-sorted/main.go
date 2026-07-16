package main

import "fmt"

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Is sorted:", isSorted(arr))
}
