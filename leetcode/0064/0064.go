package _064

// m x n
// n + m - 2 步
func minPathSum(grid [][]int) int {
	m := len(grid[0])
	n := len(grid)
	// 初始化这一步比较麻烦
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j > 0 {
				dp[0][j] = dp[0][j-1] + grid[0][j]
			}
			if i > 0 && j == 0 {
				dp[i][0] = dp[i-1][0] + grid[i][0]
			}
			if i > 0 && j > 0 {
				dp[i][j] = minValue(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[n-1][m-1]
}

func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
