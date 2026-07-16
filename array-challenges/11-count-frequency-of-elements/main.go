package main

import "fmt"

func countFrequency(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	return freq
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4}
	fmt.Println("Frequency map:", countFrequency(arr))
}
