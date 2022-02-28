package dp

func findMaxForm(strs []string, m int, n int) int {
	// m 个 0
	// n 个 1
	res := make([][]int, m+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zeroSum := 0
		oneSum := 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				zeroSum++
			} else {
				oneSum++
			}
		}

		for mIndex := m; mIndex >= zeroSum; mIndex-- {
			for nIndex := n; nIndex >= oneSum; nIndex-- {
				res[mIndex][nIndex] = getMax(res[mIndex][nIndex], res[mIndex-zeroSum][nIndex-oneSum]+1)
			}
		}
	}

	return res[m][n]
}
