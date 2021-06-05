package main

func findRepeatNumber(nums []int) int {
	a := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := a[nums[i]]
		if ok {
			return nums[i]
		}
		a[nums[i]] = 1
	}

	return -1
}
