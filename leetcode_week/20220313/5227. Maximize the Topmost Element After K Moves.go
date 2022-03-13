package _0220313

// https://leetcode-cn.com/problems/maximize-the-topmost-element-after-k-moves/solution/by-endlesscheng-vmtr/
func maximumTop(nums []int, k int) int {
	if len(nums) == 1 || k == 0 {
		if k&1 == 1 { // 奇数
			return -1
		}
		return nums[0]
	}

	if k < len(nums) {
		nums = append(nums[:k-1], nums[k])
	} else if k == len(nums) {
		nums = nums[:k-1]
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}

	return res
}
