package dp

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	left := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[left] {
			left++
		}
		if left == len(s) {
			return true
		}
	}

	return false
}

// var isSubsequence = dpIsSubsequence
func dpIsSubsequence(s string, t string) bool {
	// 最长公共子序列长度
	res := make([][]int, len(t)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(s)+1)
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else { // t 回撤 注意 s 不用回撤
				res[i][j] = res[i-1][j]
			}
		}
	}

	return len(s) == res[len(t)][len(s)]

}
