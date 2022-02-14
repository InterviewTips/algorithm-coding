package array

// 你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	slowIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slowIndex] = nums[i]
			slowIndex++
		}
	}

	return slowIndex
}
