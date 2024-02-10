package leetcode

import (
	"strings"
)

// https://leetcode-cn.com/problems/generate-parentheses/solution/gua-hao-sheng-cheng-by-leetcode-solution/
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backTrack([]string{}, 0, 0, &ans, n)
	return ans
}

func backTrack(s []string, left, right int, ans *[]string, n int) {
	if len(s) == 2*n {
		*ans = append(*ans, strings.Join(s, ""))
		return
	}
	if left < n {
		s = append(s, "(")
		backTrack(s, left+1, right, ans, n)
		s = s[:len(s)-1]
	}
	if right < left {
		s = append(s, ")")
		backTrack(s, left, right+1, ans, n)
		s = s[:len(s)-1]
	}
}
