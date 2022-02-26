package dp

func uniquePaths(m int, n int) int {
	// res[0][0] = 1
	// res[0][1] = 1
	// res[1][0] = 1
	// res[i][j] = res[i-1][j] + res[i][j-1] (i,j>=1)
	// 表示 索引位置 i,j 元素的不同路径数
	// 要么由上面走到 i,j 要么由左边走到 i,j 把两者的路径数加起来即可
	// 需要注意判断是否越界
	res := make([]int, m*n)
	// 索引 m 表示行，n 表示列
	// i * n + j
	res[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			index := i*n + j
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
