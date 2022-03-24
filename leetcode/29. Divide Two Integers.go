package leetcode

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 1 {
		return dividend
	}
	if divisor == 1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}

	sign := 1
	if (divisor^dividend)>>31&1 == 1 {
		sign = -1
	}

	var a = int64(dividend)
	var b = int64(divisor)

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	res := div(a, b)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return int(res)
	}

	// 返回负数
	return int(-res)
}

func div(a, b int64) int64 {
	if a < b {
		return 0
	}
	var count int64 = 1
	tb := b

	for tb+tb <= a {
		count += count
		tb = tb + tb
	}

	return count + div(a-tb, b)
}
