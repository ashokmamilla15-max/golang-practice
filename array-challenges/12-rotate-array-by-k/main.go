package main

import "fmt"

func rotateArray(arr []int, k int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	k = k % n
	if k < 0 {
		k += n
	}
	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
	return arr
}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Rotated array:", rotateArray(arr, 3))
}
