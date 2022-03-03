package _0220303

func addDigits(num int) int {
	for {
		num = subAddDigits(num)
		if num < 10 {
			break
		}
	}

	return num
}

func subAddDigits(num int) int {
	sum := 0
	for num/10 != 0 {
		sum += num % 10
		num = num / 10
	}
	sum += num

	return sum
}
