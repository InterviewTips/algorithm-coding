package dp

func numTrees(n int) int {
	res := make([]int, n+1)
	res[0] = 1 // 空为 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			res[i] += res[j-1] * res[i-j] // j-1 + i-j = i-1
		}
	}

	return res[n]
}
