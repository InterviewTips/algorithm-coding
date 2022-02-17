package _0220218

import "sort"

func findCenter(edges [][]int) int {
	// 3 <= n <= 105
	// edges.length == n - 1
	// 表示 len(edges) >= 2

	a := edges[0]
	b := edges[1]

	// 排序
	sort.SliceStable(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	if a[0] == b[0] {
		return a[0]
	}

	return a[1]
}
