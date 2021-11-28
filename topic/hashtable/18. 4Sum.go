package hashtable

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	// 排序先
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i <= len(nums)-4; i++ { // i <= len(nums) - 4 要简单点就直接 i < len(nums)
		num1 := nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j <= len(nums)-3; j++ { // j <= len(nums) - 3 边界条件 要简单点就 j < len(nums)
			num2 := nums[j]
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			needCalc := target - (num1 + num2)
			left := j + 1
			right := len(nums) - 1
			//log.Println(nums[left], nus[right])
			for left < right {
				if nums[left]+nums[right] == needCalc {
					res = append(res, []int{num1, num2, nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < needCalc {
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
				} else if nums[left]+nums[right] > needCalc {
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				}
			}
		}
	}

	return res
}
