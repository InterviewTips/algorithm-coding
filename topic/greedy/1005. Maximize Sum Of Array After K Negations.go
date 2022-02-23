package greedy

import (
	"log"
	"sort"
)

// 这个解法排序有问题，从小到大排序了，导致需要考虑许多边界条件
func largestSumAfterKNegations(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	index := 0

	log.Println(nums)
	for i := 0; i < k; i++ {
		if nums[index] <= 0 {
			nums[index] = -nums[index]
			if index < len(nums)-1 {
				index++
			}
		} else { // > 0
			if index > 0 {
				if nums[index-1] > nums[index] { // 上一个是正的
					nums[index] = -nums[index]
				} else {
					nums[index-1] = -nums[index-1]
				}
			} else {
				nums[index] = -nums[index]
			}
			log.Println("nums is", nums)
		}
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

// var largestSumAfterKNegations = absLargestSumAfterKNegations
func absLargestSumAfterKNegations(nums []int, k int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return getAbs(nums[i]) > getAbs(nums[j])
	})

	// 按照绝对值排序 从大到小

	log.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 { // 如果是负数 则取反
			nums[i] = -nums[i]
			k--
		}
	}

	if k&1 == 1 { // 奇数 需要取反 拿最后一个
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func getAbs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
