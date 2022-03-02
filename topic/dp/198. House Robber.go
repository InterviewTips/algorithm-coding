package dp

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := make([]int, len(nums))

	res[0] = nums[0]
	res[1] = getMax(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		res[i] = getMax(res[i-1], res[i-2]+nums[i])
	}

	return res[len(nums)-1]
}
