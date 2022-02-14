package array

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	index := len(nums) - 1 // 从末尾来
	left := 0
	right := len(nums) - 1
	for left <= right { // left == right 可能是同个索引
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare >= rightSquare {
			res[index] = leftSquare
			left++
			index--
		} else {
			res[index] = rightSquare
			right--
			index--
		}
	}

	return res
}
