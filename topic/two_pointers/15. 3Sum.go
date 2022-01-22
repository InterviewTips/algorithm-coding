package two_pointers

import "sort"

func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	// for 循环 + 快慢指针
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // nil
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 快慢指针
		slow := i + 1
		fast := len(nums) - 1
		// 0 - nums[i]
		for fast > slow {
			if nums[fast]+nums[slow]+nums[i] > 0 { // 左移
				fast--
				continue
			}
			if nums[fast]+nums[slow]+nums[i] < 0 { // 右移
				slow++
				continue
			}
			// 相等
			res = append(res, []int{nums[i], nums[slow], nums[fast]})
			// 去重
			for fast > slow && nums[fast] == nums[fast-1] {
				fast--
			}
			for fast > slow && nums[slow] == nums[slow+1] {
				slow++
			}
			// 同时收缩 处于去重之后
			fast--
			slow++
		}
	}

	return res
}
