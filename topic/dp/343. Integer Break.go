package dp

import "log"

// var integerBreak = dpIntegerBreak
func dpIntegerBreak(n int) int {
	if n == 1 {
		return 1
	}

	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	memo[2] = 1 // 其实只需要初始化 索引 2 即可 索引 0/1 没有意义
	for i := 3; i <= n; i++ {
		// 计算 integerBreak(i)
		for j := 1; j <= i-1; j++ {
			// j+(i-j)
			log.Printf("i: %v, memo[i]: %v\n", i, memo[i])
			//memo[i] = getMax3(memo[i], j*(i-j), j*memo[i-j])
			memo[i] = getMax3(memo[i], j*(i-j), j*memo[i-j])
			// 只能拆 一个 至于拆 j 还是 i-j 随意
			//memo[i] = getMax3(memo[i], (i-j) * j, (i-j) * memo[j])
			// memo[i] 需要放进来进行比较 getMax3，因为可能会进入多次 即可能拆分多次 例如 1 和 3, 2 和 2
			// 2022/02/26 15:55:39 i: 4, memo[i]: -1
			// 2022/02/26 15:55:39 i: 4, memo[i]: 3
			// 2022/02/26 15:55:39 i: 4, memo[i]: 4
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
