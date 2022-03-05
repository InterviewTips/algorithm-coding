package dp

func numDistinct(s string, t string) int {
	// 0 <= s.length, t.length <= 1000
	// 边界问题处理

	// res[i][j] 表示 [1,i] 中，出现 [1,j] 字符的次数
	res := make([][]int, len(s)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(t)+1)
	}

	// res[i][0] 可以切割出 空字符串 i > 0
	for i := 1; i < len(res); i++ {
		res[i][0] = 1
	}

	// res[0][j] 切割不了
	res[0][0] = 1

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				// 可以匹配也可以不匹配
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			} else { // s[i-1] 不匹配
				res[i][j] = res[i-1][j]
			}
		}
	}

	return res[len(s)][len(t)]
}
