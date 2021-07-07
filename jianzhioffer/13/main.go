package main

type solution struct {
	visited []bool
	rows    int
	cols    int
	maxSum  int
}

func movingCount(m int, n int, k int) int {
	// 边界条件
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	s := solution{
		visited: make([]bool, m*n),
		rows:    m,
		cols:    n,
		maxSum:  k,
	}

	// 从 0, 0 出发
	return s.getCount(0, 0)
}

func (s *solution) getCount(row, col int) int {

	count := 0
	// 边界条件
	if row >= 0 && row < s.rows && col >= 0 && col < s.cols &&
		!s.visited[row*s.cols+col] &&
		getNumSum(row, col) <= s.maxSum {
		s.visited[row*s.cols+col] = true
		count = 1 +
			s.getCount(row-1, col) + // 上
			s.getCount(row+1, col) + // 下
			s.getCount(row, col-1) + // 左
			s.getCount(row, col+1) // 右
	}

	return count

}

func getNumSum(x, y int) int {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10 // x = x / 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}

	return sum
}
