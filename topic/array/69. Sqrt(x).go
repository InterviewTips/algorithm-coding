package array

// 69. x 的平方根
//
//给你一个非负整数 x ，计算并返回 x 的 平方根 。
//
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
//
//
//示例 1：
//
//输入：x = 4
//输出：2
//
//示例 2：
//
//输入：x = 8
//输出：2
//解释：8 的平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
func mySqrt(x int) int {
	return mySqrt1(x)
}

func mySqrt1(x int) int {
	// 只需要保留整数部分，二分查找即可
	// 找到第一个 index * index 大于等于 x 的
	left := 0
	right := x // 这里边界定为 x [0, x] 如果 x 为 0 则为 [0]
	for left <= right {
		middle := left + ((right - left) >> 2)
		if middle*middle >= x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if left*left == x {
		return left
	}

	return left - 1
}

func mySqrt2(x int) int {
	// 只需要保留整数部分，二分查找即可
	// 找到 index * index 小于等于 x 的第一个索引
	left := 0
	right := x // 这里边界定为 x [0, x] 如果 x 为 0 则为 [0]
	//ans := -1
	for left <= right {
		middle := left + ((right - left) >> 2)
		if middle*middle <= x { // 哪边占据 == return 另一边 left = 则 return right
			//ans = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return right
}
