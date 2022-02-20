package _0220220

func countEven(num int) int {
	// 暴力解法
	count := 0
	for i := 1; i <= num; i++ {
		// 计算各位数字之和
		if getNumValid(i) {
			count++
		}
	}

	return count
}

func getNumValid(num int) bool {
	sum := 0
	//log.Println("num is", num)
	for num/10 != 0 {
		sum += num % 10
		num = num / 10
	}
	sum += num
	//log.Println("sum is", sum)

	return sum&1 == 0
}
