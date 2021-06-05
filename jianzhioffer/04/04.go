package _4

// 简单版
func findNumberIn2DArraySimple(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}

	return false
}

// 右上角开始找，找不到就往左边或者下边找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		// or 判断符合短路原则 一个符合后面就不会进行判断 不需要担心 panic
		return false
	}

	rows := len(matrix) - 1
	columns := len(matrix[0])
	row := 0
	column := columns - 1

	for row <= rows && column >= 0 {
		value := matrix[row][column]
		if target == value {
			return true
		}
		if target > value { // 如果目标值大于当前元素 则行移动到下一行
			row++
		} else { // target < value 如果目标值小于当前元素 则列移动到左边
			column--
		}
	}

	return false
}
