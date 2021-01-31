package _435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
}

func intervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	xEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		if start >= xEnd {
			count++
			xEnd = intervals[i][1]
		}
	}
	return count
}
