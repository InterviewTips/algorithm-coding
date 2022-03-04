package dp

func lengthOfLIS(nums []int) int {
	res := make([]int, len(nums))
	// res[i] 表示包含 j 索引下标最长上升子序列长度
	// 初始化至少是一
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}
	result := 1 // 至少也是 1
	// i > j
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				res[i] = getMax(res[i], res[j]+1)
			}
		}
		if res[i] > result {
			result = res[i]
		}
		//log.Println(res)
	}

	return result
}
