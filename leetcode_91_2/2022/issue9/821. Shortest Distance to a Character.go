package issue9

import "math"

func shortestToChar(s string, c byte) []int {
	cslice := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			// 记录 index
			cslice = append(cslice, i)
		}
	}

	var getMin func(index int) int
	getMin = func(index int) int {
		ans := math.MaxInt64
		for i := 0; i < len(cslice); i++ {
			ans = getMinValue(getAbs(cslice[i], index), ans)
		}

		return ans
	}

	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, getMin(i))
	}

	return res
}

func getMinValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getAbs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
