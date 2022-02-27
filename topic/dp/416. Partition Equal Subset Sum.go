package dp

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum&1 == 1 {
		return false
	}

	bagWeight := sum / 2

	res := make([]int, bagWeight+1)
	res[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			res[j] = getMax(res[j], res[j-nums[i]]+nums[i])
		}
	}

	return res[bagWeight] == bagWeight
}
