package leetcode

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+1) // 要求第索引行 所以要加 1
	for i := 0; i < rowIndex+1; i++ {
		res[i] = make([]int, 0)
		for j := 0; j <= i; j++ {
			if j-1 >= 0 && i-1 >= 0 && j <= i-1 {
				value := res[i-1][j-1] + res[i-1][j]
				res[i] = append(res[i], value)
			} else {
				res[i] = append(res[i], 1)
			}
		}
	}

	return res[rowIndex]
}
