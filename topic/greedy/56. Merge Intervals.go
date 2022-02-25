package greedy

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// start 要放前面
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		//log.Println(intervals[i], cur)
		if intervals[i][0] > cur[1] {
			res = append(res, cur)
			cur = intervals[i]
		} else {
			// 更新右边界
			cur[1] = getMax(intervals[i][1], cur[1])
		}
	}
	// 最后一个 cur 需要加入
	res = append(res, cur)

	return res
}
