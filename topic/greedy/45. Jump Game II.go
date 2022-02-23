package greedy

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	res := 0
	curDistance := 0
	nextDistance := 0
	for i := 0; i < len(nums); i++ {
		nextDistance = getMax(nextDistance, i+nums[i])
		if i == curDistance {
			if curDistance != len(nums)-1 {
				// 移动下
				res++
				curDistance = nextDistance
			} else {
				break
			}
		}
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
