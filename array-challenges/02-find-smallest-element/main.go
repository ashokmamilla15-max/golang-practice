package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	for _, v := range arr[1:] {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func main() {
	arr := []int{3, 5, 1, 9, 2, 8}
	fmt.Println("Smallest element:", findSmallest(arr))
}
