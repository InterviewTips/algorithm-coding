package main

import "fmt"

func removeDuplicates(nums []int) int {
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[left] != nums[right] {
			left++
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return left + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
