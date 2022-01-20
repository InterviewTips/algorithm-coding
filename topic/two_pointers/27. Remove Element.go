package two_pointers

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
