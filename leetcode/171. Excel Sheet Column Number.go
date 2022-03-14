package leetcode

func titleToNumber(columnTitle string) int {
	index := 0
	res := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		value := columnTitle[i] - 'A' + 1
		res += myPow(26, index) * int(value)
		index += 1
	}

	return res
}

func myPow(a int, b int) int {

	var subMyPow func(x int, n int) int
	subMyPow = func(x int, n int) int {
		if n == 0 {
			return 1
		}

		if n == 1 {
			return x
		}

		result := subMyPow(x, n/2)
		result *= result

		if n&1 == 1 {
			result *= x
		}

		return result
	}

	res := subMyPow(a, b)
	return res
}
