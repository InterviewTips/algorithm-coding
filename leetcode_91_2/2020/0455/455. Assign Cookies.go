package _455

import (
	"sort"
)

//输入: g = [1,2,3], s = [1,1]
//输出: 1
//输入: g = [1,2], s = [1,2,3]
//输出: 2
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	gIndex := 0
	sIndex := 0
	for gIndex < len(g) && sIndex < len(s) {
		for sIndex < len(s) && s[sIndex] < g[gIndex] {
			sIndex++
		}
		if sIndex < len(s) {
			res++
		}
		sIndex++
		gIndex++
	}
	return res
}
