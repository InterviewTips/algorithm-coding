package _0220226

func maximumDifference(nums []int) int {
	// 记录小的值
	minValue := nums[0]
	diff := -1
	// nums[i] < nums[j] 需要保证这个
	for i := 1; i < len(nums); i++ {
		if nums[i] <= minValue { // 如果不需要保证 nums[i] < nums[j] 则 nums[i] < minValue 即可
			minValue = nums[i]
			continue
		}
		diff = getMax(diff, nums[i]-minValue)
	}

	return diff
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
