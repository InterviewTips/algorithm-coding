package leetcode_hot100

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	var isInArea func(i, j int) bool
	isInArea = func(i, j int) bool {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && !visited[i][j] {
			return true
		}

		return false
	}
	var backtracking func(i, j int)
	backtracking = func(i, j int) {
		if !isInArea(i, j) {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		visited[i][j] = true
		backtracking(i+1, j)
		backtracking(i-1, j)
		backtracking(i, j+1)
		backtracking(i, j-1)
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				res += 1
				backtracking(i, j)
			}
		}
	}

	return res
}
