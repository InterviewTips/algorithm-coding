package array

// 704. 二分查找
//
//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//
//示例 1:
//
//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//示例 2:
//
//输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
func search(nums []int, target int) int {
	return search1(nums, target)
}

func search1(nums []int, target int) int {
	// 区间左闭右闭
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + ((right - left) >> 2)
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target { // target 在左边 [left, middle-1]
			right = middle - 1
		} else { // target 在右边 [middle+1, right]
			left = middle + 1
		}
	}

	return -1 // 不存在返回 -1
}

func search2(nums []int, target int) int {
	// 区间左闭右开
	left := 0
	right := len(nums)
	for left < right { // left = right 没有意义
		middle := left + ((right - left) >> 2)
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target { // [left, middle)
			right = middle
		} else { // [middle+1, right)
			left = middle + 1
		}
	}

	return -1
}
