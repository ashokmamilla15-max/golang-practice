package main

import "fmt"

//	exarr := []int{3, 5, 2, 8, 1}

func SecondLargestElement(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("array must contain at least two elements")
	}

	largest := arr[0]       //3
	secondLargest := arr[0] //3

	for _, value := range arr[1:] { //5, 2, 8, 1
		if value > largest { //3
			secondLargest = largest
			largest = value
		} else if value > secondLargest && value != largest {
			secondLargest = value
		}
	}

	if secondLargest == largest {
		return 0, fmt.Errorf("no second largest element found")
	}

	return secondLargest, nil
}

func main() {
	// Example usage
	arr := []int{3, 5, 2, 8, 1}
	secondLargest, err := SecondLargestElement(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("The second largest element is:", secondLargest)
}
