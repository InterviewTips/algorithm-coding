package _028

import "fmt"

func strStr(haystack string, needle string) int {

	// 滑动窗口
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		fmt.Printf("i: %d, haystack: %s\n", i, haystack[i:i+len(needle)])
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
