package _0220519

import "sort"

// 找出对应的中位数即可
func minMoves2(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 找出中位数
	var middle int
	if len(nums)&1 == 1 { // 奇数
		middle = nums[len(nums)/2]
	} else {
		middle = nums[len(nums)/2-1]
	}

	// 计算所有数字需要调整的次数
	count := 0
	for i := 0; i < len(nums); i++ {
		count += getAbs(nums[i], middle)
	}

	return count
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
