// Go functions can return multiple values, commonly used for a result
// paired with an error, avoiding the need for out-parameters or exceptions.
package main

import "fmt"

func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	min, max := minMax([]int{4, 2, 9, 1, 7})
	fmt.Println("min:", min, "max:", max)
}
