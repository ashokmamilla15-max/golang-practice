package main

import "fmt"

func sumAndAverage(arr []int) (int, float64) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	average := float64(sum) / float64(len(arr))
	return sum, average
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum, average := sumAndAverage(arr)
	fmt.Println("Sum:", sum, "Average:", average)
}
