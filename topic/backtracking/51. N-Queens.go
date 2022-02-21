package backtracking

import (
	"log"
)

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	visited := used{xy: make(map[[2]int]int)}
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == n {
			log.Println("path is", path)
			value := make([]string, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := 0; i < n; i++ { // y 坐标
			// if visited then continue
			log.Printf("此时 横坐标: %v, 纵坐标: %v", index, i)
			if visited.get(index, i) {
				log.Printf("已经看过 横坐标: %v, 纵坐标: %v", index, i)
				continue
			}
			path = append(path, getSubPath(i, n))
			// set visited
			visited.setXy(n, index, i)
			log.Printf("index: %v set visited is %v\n", index, visited.xy)
			backtracking(index + 1)
			// remove visited
			visited.removeXy(n, index, i)
			log.Printf("index: %v remove visited is %v\n", index, visited.xy)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	// 修改
	return res
}

type used struct {
	xy map[[2]int]int
}

func (u *used) removeXy(n int, i int, j int) {
	d := [][2]int{
		{1, 0},  // bottom
		{-1, 0}, // top
		{0, 1},  // right
		{0, -1}, // left
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	u.remove(i, j) // set self
	for index := 0; index < len(d); index++ {
		x := i
		y := j
		direction := d[index]
		for {
			x += direction[0]
			y += direction[1]
			// set x y
			if 0 <= x && x < n && 0 <= y && y < n {
				u.remove(x, y)
			} else {
				break
			}
		}
	}

	return
}

func (u *used) setXy(n int, i int, j int) {
	d := [][2]int{
		{1, 0},  // bottom
		{-1, 0}, // top
		{0, 1},  // right
		{0, -1}, // left
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	u.set(i, j) // set self
	for index := 0; index < len(d); index++ {
		x := i
		y := j
		direction := d[index]
		for {
			x += direction[0]
			y += direction[1]
			// set x y
			if 0 <= x && x < n && 0 <= y && y < n {
				u.set(x, y)
			} else {
				break
			}
		}
	}

	return
}

func (u *used) set(i, j int) {
	u.xy[[2]int{i, j}]++
}

func (u *used) get(i, j int) bool {
	return u.xy[[2]int{i, j}] > 0
}

func (u *used) remove(i, j int) {
	u.xy[[2]int{i, j}]--
}

func getSubPath(index int, n int) string {
	s := make([]byte, n)
	for i := 0; i < len(s); i++ {
		if i == index {
			s[i] = 'Q'
			continue
		}
		s[i] = '.'
	}
	return string(s)
}
