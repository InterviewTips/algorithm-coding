package _463

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i == 0 || grid[i-1][j] == 0 { //上
					res++
				}
				if i == len(grid)-1 || grid[i+1][j] == 0 { //下
					res++
				}
				if j == 0 || grid[i][j-1] == 0 { //左
					res++
				}
				if j == len(grid[0])-1 || grid[i][j+1] == 0 { //右
					res++
				}
			}
		}
	}
	return res
}
