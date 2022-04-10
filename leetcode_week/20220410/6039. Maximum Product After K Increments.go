package _0220410

import "sort"

func maximumProduct(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && k != 0 {
			nums[i] += 1
			k--
		} else {
			break
		}
	}
	if k != 0 {
		nums[0] += k
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res * nums[i] % (10e9 + 7)
	}

	return res
}
