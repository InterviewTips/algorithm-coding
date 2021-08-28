package main

func runningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i-1]
	}

	return res
}
