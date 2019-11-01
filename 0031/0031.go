package main

import "fmt"

func nextPermutation(nums []int) {
	var (
		j, i int
	)
	for i = len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 && nums[i] < nums[i+1] {
			j = i
			break
		}
	}

	// 说明单调递减
	if i == -1 {
		numsReverse(nums, 0, len(nums)-1)
	} else {
		for i = len(nums) - 1; i >= 0; i-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				numsReverse(nums, j+1, len(nums)-1)
				break
			}
		}
	}

	fmt.Println(nums)
	return
}

func numsReverse(nums []int, i int, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{1, 2, 7, 4, 3, 1})
	nextPermutation([]int{1, 2, 7, 4, 1, 3})
}
