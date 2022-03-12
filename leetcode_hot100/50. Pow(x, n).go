package leetcode_hot100

func myPow(a float64, b int) float64 {

	var subMyPow func(x float64, n int) float64
	subMyPow = func(x float64, n int) float64 {
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

	flag := false
	if b < 0 {
		flag = true
		b = -b
	}

	res := subMyPow(a, b)
	if !flag {
		return res
	}

	return 1 / res

}
