package array

//367. 有效的完全平方数
//
//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
//
//进阶：不要 使用任何内置的库函数，如  sqrt 。
//
//
//
//示例 1：
//
//输入：num = 16
//输出：true
//
//示例 2：
//
//输入：num = 14
//输出：false
func isPerfectSquare(num int) bool {
	// 结合 69 题
	left := 0
	right := num
	for left <= right {
		middle := left + ((right - left) >> 2)
		if middle*middle == num {
			return true
		}
		if middle*middle > num {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false
}
