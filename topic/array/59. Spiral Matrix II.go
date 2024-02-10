package array

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 示例 1：
//
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
// 示例 2：
//
// 输入：n = 1
// 输出：[[1]]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/spiral-matrix-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func generateMatrix(n int) [][]int {
	loop := n / 2 // 画几圈
	mid := n / 2  //
	// 初始化 res
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	count := 1 // 从 1 开始画
	startX, startY := 0, 0
	offset := 1 // 边界问题
	// 开始画圈
	for loop > 0 {
		// 遵循左闭右开
		i := startX
		j := startY
		// 上 循环出来之后 j 是一行中最后的索引
		for ; j < startY+n-offset; j++ {
			res[i][j] = count
			count++
		}
		// 右
		for ; i < startX+n-offset; i++ {
			res[i][j] = count
			count++
		}
		// 下
		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}
		// 左
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}

		startX++
		startY++
		offset += 2 // 画一圈之后两旁的数字都被画上了
		loop--
	}

	if n%2 == 1 { // 奇数
		res[mid][mid] = count
	}

	return res
}
