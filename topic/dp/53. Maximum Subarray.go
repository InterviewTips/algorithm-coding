package dp

func maxSubArray(nums []int) int {
	// res[j] 表示包含 nums[j] 的最大的子数组和
	res := make([]int, len(nums))

	res[0] = nums[0]
	result := res[0]
	for i := 1; i < len(res); i++ {
		res[i] = getMax(res[i-1]+nums[i], nums[i])
		if res[i] > result {
			result = res[i]
		}
	}

	return result
}
