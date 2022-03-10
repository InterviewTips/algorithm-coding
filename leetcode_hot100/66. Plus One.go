package leetcode_hot100

func plusOne(digits []int) []int {
	left := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			value := digits[i] + 1
			digits[i] = value % 10
			left = value / 10
			continue
		}
		value := digits[i] + left
		digits[i] = value % 10
		left = value / 10
	}
	if left != 0 {
		res := make([]int, 0)
		res = append(res, left)
		res = append(res, digits...)
		return res
	}
	return digits
}
