package backtracking

func totalNQueens(n int) int {
	count := 0
	path := make([]string, 0)
	visited := used{xy: make(map[[2]int]int)}
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == n {
			count++
			return
		}

		for i := 0; i < n; i++ { // y 坐标
			// if visited then continue
			// log.Printf("此时 横坐标: %v, 纵坐标: %v", index, i)
			if visited.get(index, i) {
				// log.Printf("已经看过 横坐标: %v, 纵坐标: %v", index, i)
				continue
			}
			path = append(path, getSubPath(i, n))
			// set visited
			visited.setXy(n, index, i)
			// log.Printf("index: %v set visited is %v\n", index, visited.xy)
			backtracking(index + 1)
			// remove visited
			visited.removeXy(n, index, i)
			// log.Printf("index: %v remove visited is %v\n", index, visited.xy)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	// 修改
	return count
}
