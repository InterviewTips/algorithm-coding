package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	//log.Println(points)

	count := 1 // 表示非重复区间个数
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < cur[1] {
			// 关键: 需要保存最小右边界
			cur[1] = getMin(cur[1], intervals[i][1])
			continue
		} else {
			count++
			cur = intervals[i]
		}
	}

	return len(intervals) - count // 需要移除的个数
}
