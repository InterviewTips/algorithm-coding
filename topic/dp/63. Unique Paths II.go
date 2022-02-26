package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// res[0][0] 不一定等于 1
	// 官方示例中给了一个 [[1]] 导致无路可走
	// 遇到石头 则为 0 因为不可能走到石头上面
	// res[i][j] = res[i-1][j] + res[i][j-1] (i,j>=1)
	// 表示 索引位置 i,j 元素的不同路径数
	// 要么由上面走到 i,j 要么由左边走到 i,j 把两者的路径数加起来即可
	// 需要注意判断是否越界
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([]int, m*n)
	// 索引 m 表示行，n 表示列
	// i * n + j
	res[0] = 1
	if obstacleGrid[0][0] == 1 { // 特殊处理下
		res[0] = 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			index := i*n + j
			if obstacleGrid[i][j] == 1 { // 是石头 不用往下 直接为 0
				res[index] = 0
				continue
			}
			if i-1 >= 0 {
				res[index] += res[(i-1)*n+j]
			}
			if j-1 >= 0 {
				res[index] += res[i*n+j-1]
			}
		}
	}

	return res[m*n-1]
}
