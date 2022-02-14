package array

import "math"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 双指针
	left := 0
	sum := 0
	subLength := 0
	res := math.MaxInt64
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			subLength = right - left + 1
			if subLength < res {
				res = subLength
			}
			// 移除 left
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt64 {
		res = 0
	}

	return res
}
