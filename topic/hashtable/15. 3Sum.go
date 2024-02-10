package hashtable

import (
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	// 排序先
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//log.Println(nums)
	//判断去重
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if nums[left]+nums[right] > target {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}

	return res

}
