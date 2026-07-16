package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			merged = append(merged, curr)
		}
	}
	return merged
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("Merged intervals:", mergeIntervals(intervals))
}
