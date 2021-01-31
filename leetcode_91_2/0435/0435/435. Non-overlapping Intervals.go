package _435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	xEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		if start < xEnd {
			res++
		} else {
			xEnd = intervals[i][1]
		}
	}
	return res
}
