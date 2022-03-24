package _0220324

func imageSmoother(img [][]int) [][]int {
	// 上下左右
	d := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}
	res := make([][]int, len(img))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(img[0]))
	}

	calc := func(num int, i, j int) int {
		count := 1
		sum := num
		for index := 0; index < len(d); index++ {
			dI := d[index][0] + i
			dJ := d[index][1] + j
			if dI >= 0 && dI < len(img) && dJ >= 0 && dJ < len(img[0]) {
				count++
				sum += img[dI][dJ]
			}
		}

		return sum / count
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			res[i][j] = calc(img[i][j], i, j)
		}
	}

	return res
}
