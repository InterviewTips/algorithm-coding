package array

func generateMatrix(n int) [][]int {
	loop := n / 2
	mid := n / 2
	count := 1 // 生成值
	startx := 0
	starty := 0
	offset := 1 // 每一圈控制的遍历长度

	// 初始化 res
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	for ; loop != 0; loop-- {
		i := startx // x 坐标 横向
		j := starty // y 坐标 纵向

		// 上边 左 -> 右
		for j = starty; j < starty+n-offset; j++ {
			res[i][j] = count
			count++
		}
		// 右边 上 -> 下
		for i = startx; i < startx+n-offset; i++ {
			res[i][j] = count
			count++
		}
		// 下边 右 -> 左 j != starty
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		// 左边 下 -> 上 i != startx
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}

		// 移动下一层的开始坐标
		startx++
		starty++

		// offset 关键
		offset += 2 // 左边都是两行

	}

	// 中间点关键
	if n&1 == 1 {
		res[mid][mid] = count
	}

	return res
}
