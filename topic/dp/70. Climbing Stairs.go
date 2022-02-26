package dp

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// res[i] 表示爬 i 层台阶有多少种方式
	// res[i] = res[i-1] + res[i-2]
	// 最后只爬 1 层，则有 res[i-1] 种方式
	// 最后只爬 2 层，则有 res[i-2] 种方式
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	for i := 3; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
