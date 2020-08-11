package _055

// 算出长度
// 拿到第一个位置的数字
func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
