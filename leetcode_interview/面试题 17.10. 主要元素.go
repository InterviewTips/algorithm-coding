package leetcode_interview

// https://leetcode.cn/problems/find-majority-element-lcci/solution/zhu-yao-yuan-su-by-leetcode-solution-xr1p/
// Boyer-Moore
func majorityElement(nums []int) int {
	count := 0
	candidate := -1
	for i := 0; i < len(nums); i++ {
		item := nums[i]
		if count == 0 {
			candidate = item
		}
		if item == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for i := 0; i < len(nums); i++ {
		item := nums[i]
		if item == candidate {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	}

	return -1
}
