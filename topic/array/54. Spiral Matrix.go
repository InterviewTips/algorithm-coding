package array

// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, m*n)

	var loop int
	var mid int
	var middle bool
	if m > n {
		loop = n / 2 // 画几圈
		mid = n / 2  //
		if n%2 == 1 {
			middle = true
		}
	} else {
		loop = m / 2 // 画几圈
		mid = m / 2  //
		if m%2 == 1 {
			middle = true
		}
	}

	index := 0
	startX, startY := 0, 0
	offset := 1 // 边界问题
	// 开始画圈
	for loop > 0 {
		// 遵循左闭右开
		i := startX
		j := startY
		// 上 循环出来之后 j 是一行中最后的索引
		for ; j < startY+n-offset; j++ {
			res[index] = matrix[i][j]
			index++
		}
		// 右
		for ; i < startX+m-offset; i++ {
			res[index] = matrix[i][j]
			index++
		}
		// 下
		for ; j > startY; j-- {
			res[index] = matrix[i][j]
			index++
		}
		// 左
		for ; i > startX; i-- {
			res[index] = matrix[i][j]
			index++
		}

		startX++
		startY++
		offset += 2 // 画一圈之后两旁的数字都被画上了
		loop--
	}

	if middle { // 奇数
		// 从 mid 再划一遍
		i := startX
		j := startY
		if m < n { // 横排
			for ; j <= startY+n-offset; j++ {
				res[index] = matrix[mid][j]
				index++
			}
		}
		if m > n { // 竖排
			for ; i <= startX+m-offset; i++ {
				res[index] = matrix[i][mid]
				index++
			}
		}
		if m == n {
			res[index] = matrix[mid][mid]
			index++
		}
	}

	return res
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, m*n)

	top := 0
	left := 0
	bottom := m - 1
	right := n - 1
	index := 0
	// 开始画圈
	for index <= m*n-1 {
		// 注意：for 循环中需要进行 && index <= m *n - 1 此限定条件 不然有可能会画多
		for i := left; i <= right && index <= m*n-1; i++ {
			res[index] = matrix[top][i]
			index++
		}
		top++
		for i := top; i <= bottom && index <= m*n-1; i++ {
			res[index] = matrix[i][right]
			index++
		}
		right--
		for i := right; i >= left && index <= m*n-1; i-- {
			res[index] = matrix[bottom][i]
			index++
		}
		bottom--
		for i := bottom; i >= top && index <= m*n-1; i-- {
			res[index] = matrix[i][left]
			index++
		}
		left++
	}

	return res
}
