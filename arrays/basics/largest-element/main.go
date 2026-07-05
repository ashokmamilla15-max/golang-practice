package main

import "fmt"

func LargestElement(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	largest := arr[0]
	for _, value := range arr[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest, nil
}

func main() {
	// Example usage
	arr := []int{3, 5, 2, 8, 1}
	largest, err := LargestElement(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	println("The largest element is:", largest)
}
