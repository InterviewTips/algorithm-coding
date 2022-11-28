package _0221127

func check(nums []int) bool {
	n := len(nums)
	x := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] { // 后一个变小
			x = i
			break
		}
	}
	if x == 0 { // 要么是长度为1，要么递增
		return true
	}
	for i := x + 1; i < n; i++ {
		if nums[i] < nums[i-1] { // 后一个再次变小，直接返回 false
			return false
		}
	}

	return nums[0] >= nums[n-1] // 最后的判定条件，首元素 >= 尾元素
}
