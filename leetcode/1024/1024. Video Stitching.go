package _024

import (
	"sort"
)

func videoStitching(clips [][]int, T int) int {
	// sort
	sort.SliceStable(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0] || (clips[i][0] == clips[j][0] && clips[i][1] > clips[j][1])
	})
	//fmt.Printf("%v\n", clips)
	if clips[0][0] != 0 {
		return -1
	}
	start := clips[0]
	count := 1
	for i := 1; i < len(clips); i++ {
		if (clips[i][0] <= start[1]) &&
			(clips[i][1] > start[1]) {
			count += 1
			start = clips[i]
			if clips[i][1] >= T {
				return count
			}
		}
	}
	return -1
}
