package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var (
		numsLen = len(nums)
		res     [][]int
		i, j    int
	)
	res = make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)

	for i = 0; i < numsLen-3; i++ {
		// 跳过重复的值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j = i + 1; j < numsLen-2; j++ {
			// 不加 j > i+1 的话会漏掉 nums[i] 与 nums[j] 数值相等的情况
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := numsLen - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				} else {

					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				}

			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
	fmt.Println(fourSum([]int{-1, -5, -5, -3, 2, 5, 0, 4}, -7))
}
