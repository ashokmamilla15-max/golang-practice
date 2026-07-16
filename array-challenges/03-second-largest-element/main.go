package main

import (
	"fmt"
	"math"
)

func secondLargest(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("array must contain at least two elements")
	}
	first, second := math.MinInt, math.MinInt
	for _, v := range arr {
		switch {
		case v > first:
			second = first
			first = v
		case v > second && v != first:
			second = v
		}
	}
	if second == math.MinInt {
		return 0, fmt.Errorf("no distinct second largest element")
	}
	return second, nil
}

func main() {
	arr := []int{3, 5, 1, 9, 2, 8}
	second, err := secondLargest(arr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Second largest element:", second)
}
