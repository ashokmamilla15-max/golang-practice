package main

import "fmt"

func moveZeroes(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println("Array with zeroes moved:", moveZeroes(arr))
}
