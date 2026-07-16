package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]int, 0)
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}
	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Spiral order:", spiralOrder(matrix))
}
