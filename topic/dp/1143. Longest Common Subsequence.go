package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	// res[i][j] 索引 [0,i-1] 的字符串和 [0, j-1] 的字符串
	res := make([][]int, len(text1)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = getMax(res[i-1][j], res[i][j-1])
			}
		}
	}

	return res[len(text1)][len(text2)]
}
