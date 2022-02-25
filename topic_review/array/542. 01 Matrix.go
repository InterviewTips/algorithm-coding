package array

func updateMatrix(mat [][]int) [][]int {
	// 计算所有 1 到 所有 0 的距离
	// 取出最小值
	// 如果 1 被堵住了 说明是他周围的几个 1 到 0 的距离加 1 的最小值
	row := len(mat)
	col := len(mat[0])
	res := make([][]int, len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(mat[0]))
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = 20000
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if j-1 >= 0 {
				res[i][j] = getMin(res[i][j], res[i][j-1]+1)
			}
			if i-1 >= 0 {
				res[i][j] = getMin(res[i][j], res[i-1][j]+1)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i+1 < row {
				res[i][j] = getMin(res[i][j], res[i+1][j]+1)
			}
			if j+1 < col {
				res[i][j] = getMin(res[i][j], res[i][j+1]+1)
			}
		}
	}

	return res
}

func getMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
