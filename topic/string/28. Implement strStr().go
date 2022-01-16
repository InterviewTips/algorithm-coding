package string

func strStr(haystack string, needle string) int {
	if len(needle) == 0 { // 边界条件
		return 0
	}
	next := getNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] { // 不相等需要不断回退而不是回退一次 所以使用 for
			j = next[j-1] // 回退到上一个前缀相同的地方的下一个索引
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1 // 此时 i 还没有 ++
		}
	}

	return -1
}

//getNext 前缀表构建
func getNext(needle string) []int {
	next := make([]int, len(needle))
	j := 0
	next[0] = 0
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	return next
}
