package dp

func longestPalindromeSubseq(s string) int {
	res := make([][]int, len(s))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(s))
	}

	// init
	for i := 0; i < len(s); i++ {
		res[i][i] = 1
	}

	// j > i
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				res[i][j] = res[i+1][j-1] + 2
			} else {
				res[i][j] = getMax(res[i][j-1], res[i+1][j])
			}
		}
	}

	return res[0][len(s)-1]
}
