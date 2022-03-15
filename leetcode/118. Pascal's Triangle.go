package leetcode

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
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

	return res
}
