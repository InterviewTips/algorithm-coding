package _435

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - mostIntervals(intervals)
}

func mostIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// sort
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Printf("intervals is %v\n", intervals)
	count := 1
	end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		if start >= end {
			count++
			end = intervals[i][1]
		}
	}
	fmt.Printf("count is %d\n", count)
	return count
}
