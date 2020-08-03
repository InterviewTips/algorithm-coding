package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var (
		res     [][]int
		numsLen int
		i       int
	)

	res = make([][]int, 0)
	numsLen = len(nums)
	if numsLen < 3 || nums == nil {
		return res
	}
	// 排序
	sort.Ints(nums)
	// 循环计算
	for i = 0; i < numsLen-2; i++ {
		// 排序过后，如果 num[i] > 0 后面的肯定也 > 0 直接结束循环
		if nums[i] > 0 {
			break
		}
		// 如果 num[i] 和 前一个数一样 则会产生重复结果 去除
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := numsLen - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 避免产生同样的结果 跳过
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 避免产生同样的结果 跳过
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				//	如果 sum < 0 则 l 需要往右边移动 数字才会增大 sum 才有可能等于 0
				l++
			} else {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			}

		}
	}

	return res

}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
}
