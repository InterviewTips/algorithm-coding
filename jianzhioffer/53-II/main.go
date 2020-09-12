package main

/**
一个长度为 n-1 的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围 0～n-1 之内。
在范围 0～n-1 内的 n 个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 直观思维
func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 2 {
			return nums[i] - 1
		}
	}
	return len(nums)
}

// nums[i] == i 遍历思路
func missingNumber1(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

// 二分查找
func missingNumber2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		m := left + (right-left)/2
		if nums[m] == m { // 说明在右半部分
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return left
}
