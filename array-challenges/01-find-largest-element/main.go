package main

import "fmt"

func findLargest(arr []int) int {
	largest := arr[0]
	for _, v := range arr[1:] {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	arr := []int{3, 5, 1, 9, 2, 8}
	fmt.Println("Largest element:", findLargest(arr))
}
