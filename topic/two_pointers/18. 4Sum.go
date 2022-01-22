package two_pointers

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			slow := j + 1
			fast := len(nums) - 1
			//log.Println(nums)
			//log.Println(i, j, slow, fast)
			for fast > slow {
				if nums[fast]+nums[slow]+nums[i]+nums[j] > target {
					fast--
					continue
				}
				if nums[fast]+nums[slow]+nums[i]+nums[j] < target {
					slow++
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[slow], nums[fast]})
				for fast > slow && nums[fast] == nums[fast-1] {
					fast--
				}
				for fast > slow && nums[slow] == nums[slow+1] {
					slow++
				}
				fast--
				slow++
			}
		}
	}

	return res
}
