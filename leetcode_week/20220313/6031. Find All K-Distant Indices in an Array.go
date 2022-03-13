package _0220313

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	// 找出 key 的索引
	keyIndex := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			keyIndex = append(keyIndex, i)
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// i - j <= k
		for j := 0; j < len(keyIndex); j++ {
			if getAbs(i, keyIndex[j]) <= k {
				res = append(res, i)
				break
			}
		}
	}

	// 排序
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
