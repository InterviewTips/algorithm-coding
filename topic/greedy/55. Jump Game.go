package greedy

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	max := 0
	for i := 0; i <= max; i++ {
		if max >= len(nums)-1 {
			return true
		}
		if value := nums[i] + i; value > max {
			max = value
		}
	}

	return false
}
