package array

func updateMatrix(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				continue
			}
			// todo:
			mat[i][j] = 0
		}
	}

	return mat
}