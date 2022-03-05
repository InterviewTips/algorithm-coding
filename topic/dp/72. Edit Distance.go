package dp

import "log"

// var minDistance = minDistance2
func minDistance2(word1 string, word2 string) int {
	res := make([][]int, len(word1)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(word2)+1)
	}

	// 初始化
	res[0][0] = 0
	for i := 1; i < len(res); i++ {
		res[i][0] = i
	}
	for j := 1; j < len(res[0]); j++ {
		res[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = getMin(res[i-1][j-1]+1, getMin(res[i-1][j]+1, res[i][j-1]+1))
			}
		}
	}
	log.Println(res)

	return res[len(word1)][len(word2)]
}
