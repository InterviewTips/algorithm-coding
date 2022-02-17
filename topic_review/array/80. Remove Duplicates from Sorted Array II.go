package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slowIndex := 0
	count := 1 // 记录数字出现的次数
	for rightIndex := 1; rightIndex < len(nums); rightIndex++ {
		if nums[rightIndex] != nums[slowIndex] {
			slowIndex++
			nums[slowIndex] = nums[rightIndex]
			count = 1
			continue
		}
		if nums[rightIndex] == nums[slowIndex] && count == 1 {
			slowIndex++
			nums[slowIndex] = nums[rightIndex]
			count = 2
		}
	}

	return slowIndex + 1
}
