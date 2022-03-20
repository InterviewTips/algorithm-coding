package _0220320

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	// 12 个，2的n次方子集
	// 需要注意有可能 NumArrows 填充不完 所以要填充下
	var res []int
	maxScore := -1
	for i := 0; i < 1<<12; i++ { // 枚举所有集合
		score, arrows, bobArrows := 0, 0, [12]int{}
		for j := 0; j < 12; j++ {
			if i&(1<<j) > 0 { // bob 得分 这里是 > 0 而不是 == 1 如果要等于 == 1 的判断条件为 if i>>j&1 ==1
				score += j                        // 是 j 而不是 j+1 得分是 0-11 而不是 1-12
				bobArrows[j] = aliceArrows[j] + 1 // 节省箭数量
				arrows += bobArrows[j]            // 使用的箭的数量
			}
		}
		if arrows > numArrows { // 超过 numArrows
			continue
		}
		if score > maxScore {
			maxScore = score
			bobArrows[0] = numArrows - arrows
			res = bobArrows[:]
		}
	}

	return res
}
