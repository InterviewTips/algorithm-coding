package main

// 思路：暴力回溯
// 从 board[0][0] 开始遍历，注意边界条件

type solution struct {
	pathLength int
	word       string
	board      [][]byte
	rows       int
	cols       int
	visited    []bool
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	s := solution{
		pathLength: 0,
		word:       word,
		board:      board,
		rows:       len(board),
		cols:       len(board[0]),
		visited:    make([]bool, len(board)*len(board[0])),
	}

	// 遍历
	for row := 0; row < s.rows; row++ {
		for col := 0; col < s.cols; col++ {
			if s.isPath(row, col) {
				return true
			}
		}
	}

	// 无解
	return false
}

func (s *solution) isPath(row, col int) bool {
	if len(s.word) == s.pathLength { // 走到了最后一个字符
		return true
	}

	hasPath := false
	// 边界条件 rows cols 相等 没被走过
	if row >= 0 && row < s.rows && col >= 0 && col < s.cols &&
		s.board[row][col] == s.word[s.pathLength] &&
		!s.visited[row*s.cols+col] {

		s.visited[row*s.cols+col] = true // 被走过了
		s.pathLength++                   // 走了一步

		// 上下左右
		hasPath = s.isPath(row-1, col) || s.isPath(row+1, col) ||
			s.isPath(row, col-1) || s.isPath(row, col+1)

		if !hasPath { // 重置
			s.visited[row*s.cols+col] = false
			s.pathLength--
		}
	}

	return hasPath

}
