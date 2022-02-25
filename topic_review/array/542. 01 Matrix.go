package array

import "log"

func updateMatrix(mat [][]int) [][]int {
	// 计算所有 1 到 所有 0 的距离
	// 取出最小值
	// 如果 1 被堵住了 说明是他周围的几个 1 到 0 的距离加 1 的最小值

	visited := make([][]bool, len(mat))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(mat[0]))
	}

	res := make([][]int, len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(mat[0]))
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = -1
		}
	}

	var subUpdateMatrix func(i, j int) int
	subUpdateMatrix = func(i, j int) int {
		// set x y
		if mat[i][j] == 0 {
			return 0
		}
		if res[i][j] != -1 {
			return res[i][j]
		}

		d := [][2]int{
			{1, 0},  // bottom
			{-1, 0}, // top
			{0, 1},  // right
			{0, -1}, // left
		}

		//res := make([]int, 0)

		visited[i][j] = true
		minValue := 20000
		for index := 0; index < len(d); index++ {
			x := i
			y := j
			direction := d[index]
			x += direction[0]
			y += direction[1]
			// set x y
			if 0 <= x && x < len(mat) && 0 <= y && y < len(mat[0]) && !visited[x][y] {
				//if mat[x][y] == 0 { // 如果周围是 0，则返回 1
				//	return 1
				//}
				if res[x][y] != -1 {
					log.Printf("res exists, i: %v, j: %v, x: %v, y: %v, value: %v", i, j, x, y, res[x][y])
					value := res[x][y]
					if value < minValue {
						minValue = value
					}
					continue
				}
				visited[x][y] = true
				value := subUpdateMatrix(x, y)
				if value == 0 {
					visited[x][y] = false
					minValue = value
					break
				}
				log.Printf("i: %v, j: %v, x: %v, y: %v, value: %v", i, j, x, y, value)
				if value < minValue {
					minValue = value
				}
				visited[x][y] = false
			}
		}

		visited[i][j] = false
		//res[i][j] = minValue + 1
		// 周围都是 1 说明需要取上下左右四个方向+1 的最小值
		return minValue + 1

	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			}
			res[i][j] = subUpdateMatrix(i, j)
		}
	}

	return res
}
