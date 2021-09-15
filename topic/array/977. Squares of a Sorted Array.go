package array

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	pos := len(nums) - 1
	res := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[pos] = nums[right] * nums[right]
			right--
		} else {
			res[pos] = nums[left] * nums[left]
			left++
		}
		pos--
	}

	return res
}
