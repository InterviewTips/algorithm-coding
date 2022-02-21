package backtracking

func solveSudoku(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for i := 0; i < len(board); i++ { // 行
			for j := 0; j < len(board[0]); j++ { // 列
				if board[i][j] != '.' {
					continue
				}
				for char := '1'; char <= '9'; char++ { // (i,j) 放 k 是否合适
					if isValidNum(i, j, byte(char), board) {
						board[i][j] = byte(char)
						if backtracking() { // 找到一组合适的数独则返回
							return true
						}
						board[i][j] = '.'
					}
				}
				// 9 个数字都没有办法填充这一个格
				return false
			}
		}

		return true
	}

	backtracking()
	return
}

func isValidNum(row int, col int, char byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == char {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if board[j][col] == char {
			return false
		}
	}

	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == char {
				return false
			}
		}
	}

	return true
}
