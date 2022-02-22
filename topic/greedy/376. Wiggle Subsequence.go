package greedy

func wiggleMaxLength(nums []int) int {
	res := 1
	preDiff := 0
	curDiff := 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (preDiff <= 0 && curDiff > 0) ||
			(preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}

	return res
}
