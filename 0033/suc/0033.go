package main

import "fmt"

func search(nums []int, target int) int {

	var (
		mid int
		l   = 0
		r   = len(nums) - 1
	)

	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] <= nums[r] {
			// 右边是有序的
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// 左边是有序的
			if nums[l] <= target && target < nums[mid]   {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1

}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2},3))
}
