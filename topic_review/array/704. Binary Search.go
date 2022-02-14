package array

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 左闭右闭
	for left <= right { // [left, right] 是有意义的
		middle := left + (right-left)/2 // 防止溢出 如果直接 right+left 的话
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > target { // 更新 right
			right = middle - 1
		} else if nums[middle] < target { // 更新 left
			left = middle + 1
		}
	}

	return -1
}
