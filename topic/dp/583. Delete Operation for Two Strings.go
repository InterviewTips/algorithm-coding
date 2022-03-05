package dp

func minDistance(word1 string, word2 string) int {
	// 求出最长公共子序列长度
	res := make([][]int, len(word1)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = getMax(res[i-1][j], res[i][j-1])
			}
		}
	}

	value := res[len(word1)][len(word2)]

	return len(word1) - value + len(word2) - value
}
