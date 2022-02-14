package _0220214

//func singleNonDuplicate(nums []int) int {
//	// O(n) 时间复杂度 O(n)空间复杂度
//	data := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		data[nums[i]]++
//		if data[nums[i]] == 2 {
//			delete(data, nums[i])
//		}
//	}
//
//	for k, v := range data {
//		if v == 1 {
//			return k
//		}
//	}
//
//	// 正常不会走到这里来
//	return -1
//}

// 0 <= nums[i] <= 105
// O(n) 时间复杂度
// O(1) 空间复杂度
//func singleNonDuplicate(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	var pre = -1
//	// 有序
//	for i := 0; i < len(nums); i++ {
//		pre = nums[i]
//		i++
//		if i < len(nums)-1 {
//			if nums[i] != pre {
//				return pre
//			} // 此时已经找到
//		}
//	}
//
//	return pre
//}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		middle -= middle & 1
		if middle+1 > right { // 不能比 right 大 不然可能 out of range
			break
		}
		if nums[middle] == nums[middle+1] {
			left = middle + 2
		} else {
			right = middle - 1
		}
	}

	return nums[left]
}
