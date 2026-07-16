package main

import "fmt"

func longestConsecutive(arr []int) int {
	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}
	longest := 0
	for v := range set {
		if !set[v-1] {
			length := 1
			for set[v+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Longest consecutive sequence:", longestConsecutive(arr))
}
