package greedy

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	//log.Println(points)

	count := 1
	cur := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= cur[1] {
			// 关键: 需要保存最小右边界
			cur[1] = getMin(cur[1], points[i][1])
			continue
		} else {
			count++
			cur = points[i]
		}
	}

	return count
}

func getMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
