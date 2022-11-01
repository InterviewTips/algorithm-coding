package issue9

func addToArrayForm(num []int, k int) []int {
	res := make([]int, 0)
	// 逐位相加
	for i := len(num) - 1; i >= 0; i-- {
		sum := num[i] + k%10
		k /= 10
		if sum >= 10 {
			k++       // 进位
			sum -= 10 // 减掉 10
		}
		res = append(res, sum)
	}
	for k != 0 { // 如果有剩余的接着加
		res = append(res, k%10)
		k /= 10
	}

	// reverse
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
