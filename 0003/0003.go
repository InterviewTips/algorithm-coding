package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var ans int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right] // s[0:0] 为空
	// 遍历所有 s[:right] 的情况
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			// 如果存在 则下一次从 left + index + 1 开始
			left += index + 1
		}
		s1 = s[left : right+1] // 记录下 s1 的长度并与 ans 作比较
		if len(s1) > ans {
			ans = len(s1)
		}
	}

	return ans
}

func main() {
	println(lengthOfLongestSubstring("hello"))
}
