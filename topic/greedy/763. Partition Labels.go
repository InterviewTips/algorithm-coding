package greedy

func partitionLabels(s string) []int {
	// 记录每个字母出现下标出现的最终位置
	data := make(map[byte]int)
	res := make([]int, 0)
	for i := 0; i < len(s); i++ {
		data[s[i]] = i
	}

	left := 0
	right := 0
	for i := 0; i < len(s); i++ {
		right = getMax(right, data[s[i]])
		if i == right {
			res = append(res, right-left+1) // 插入的元素是长度 不是索引
			left = i + 1
		}
	}

	return res
}
