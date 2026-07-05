package main

import "fmt"

func SmallestElement(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	smallest := arr[0]
	for _, value := range arr[1:] {
		if value < smallest {
			smallest = value
		}
	}
	return smallest, nil
}

func main() {
	// Example usage
	arr := []int{3, 5, 2, 8, 1}
	smallest, err := SmallestElement(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	println("The smallest element is:", smallest)
}
