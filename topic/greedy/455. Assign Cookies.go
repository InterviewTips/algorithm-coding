package greedy

import "sort"

// findContentChildren
// @Description: 分发饼干
// @param g 小孩数组
// @param s 饼干数组
// @return int 满足数量
func findContentChildren(g []int, s []int) int {
	count := 0
	sort.SliceStable(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	gIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= g[gIndex] {
			gIndex++
			count++
		}
		if gIndex >= len(g) {
			break
		}
	}

	return count
}
