package main

import "fmt"

func searchRange(nums []int, target int) []int {
	// 时间复杂度 O(log n) + O(log n) = O(log n)
	return []int{leftSearch(nums, target), rightSearch(nums, target)}
}

func leftSearch(nums []int, target int) int {
	var (
		l = 0
		r = len(nums) - 1
	)
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			//	左边界搜索
			r = mid - 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l == len(nums) {
		return -1
	} else {
		if nums[l] == target {
			return l
		} else {
			return -1
		}
	}
}

func rightSearch(nums []int, target int) int {
	var (
		l = 0
		r = len(nums) - 1
	)
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			// 右边界搜索
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		// 说明没有出现过 target == nums[mid] 也没有 target > nums[mid]
		return -1
	} else {
		if nums[l-1] == target {
			return l - 1
		} else {
			return -1
		}
	}
}

func main() {
	//fmt.Println(searchRange([]int{1, 2, 2, 3, 4}, 2))
	fmt.Println(searchRange([]int{1, 2, 2, 3, 4}, 3))
}
