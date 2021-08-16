package main

func myPow(base float64, exponent int) float64 {
	if base == 0 && exponent < 0 { // 无意义
		return 0
	}

	absExponent := getAbsExponent(exponent)

	res := getResult(base, absExponent)
	if exponent < 0 {
		res = 1.0 / res
	}

	return res
}

func getResult(base float64, exponent int) float64 { // 递归

	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := getResult(base, exponent>>1)
	result *= result

	if exponent&0x1 == 1 {
		result *= base
	}

	return result
}

func getAbsExponent(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}
