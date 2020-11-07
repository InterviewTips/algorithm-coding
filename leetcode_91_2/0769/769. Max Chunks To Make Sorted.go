package _769

//注：序列为 [0..n-1]
func maxChunksToSorted(arr []int) int {
	max := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		max = maxValue(max, arr[i])
		if max == i { // 关键
			count++
		}
	}
	return count
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
