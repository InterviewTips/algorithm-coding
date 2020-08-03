package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var (
		nMin    int
		numsLen = len(nums)
		i       int
		res     int
	)

	numsLen = len(nums)
	nMin = mathAbs(nums[0]+nums[1]+nums[2], target)

	sort.Ints(nums)

	for i = 0; i < numsLen; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := numsLen - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			} else if sum < target {
				// a+b+c < target
				if mathAbs(sum, target) <= nMin {
					nMin = mathAbs(sum, target)
					res = sum
				}
				l++
			} else {
				// a+b+c > target
				if mathAbs(sum, target) <= nMin {
					nMin = mathAbs(sum, target)
					res = sum
				}
				r--
			}

		}
	}

	return res
}

// 获取绝对值
func mathAbs(x, y int) int {
	if (x - y) > 0 {
		return x - y
	}
	return y - x
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))

}
