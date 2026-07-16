package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{4, 2, 7, 1, 9}
	fmt.Println("Index of 7:", linearSearch(arr, 7))
}
