package _066

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{}
	}
	for i := n - 1; i >= 0; i-- {
		if digits[i]+1 == 10 { // 进位
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	// 你可以假设除了整数 0 之外，这个整数不会以零开头。
	// 走到这里说明需要补 1
	digits = append([]int{1}, digits...)
	return digits
}
