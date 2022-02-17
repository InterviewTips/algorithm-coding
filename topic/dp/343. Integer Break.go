package dp

// var integerBreak = dpIntegerBreak
func dpIntegerBreak(n int) int {
	if n == 1 {
		return 1
	}

	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 计算 integerBreak(i)
		for j := 1; j <= i-1; j++ {
			// j+(i-j)
			memo[i] = getMax3(memo[i], j*(i-j), j*memo[i-j])
		}
	}

	return memo[n]
}

func integerBreak(n int) int {
	if n == 1 {
		return 1
	}

	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	return subInterBreak(n, memo)
}

// 将其拆分为至少2部分
func subInterBreak(n int, memo []int) int {
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}
	// n>=2
	// 回溯
	res := -1
	for i := 1; i <= n-1; i++ {
		// 这里可以不继续分，所以需要比较 i*(n-i)
		res = getMax3(res, i*(n-i), i*subInterBreak(n-i, memo))
	}

	memo[n] = res

	return res
}

func getMax3(a, b, c int) int {
	return getMax(a, getMax(b, c))
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
