package leetcode

func findNumbers(nums []int) int {
	count := 0 // 位数为偶数的个数
	for i := 0; i < len(nums); i++ {
		item := nums[i]
		if judge(item) {
			// fmt.Println(item)
			count++
		}
	}

	return count
}

func judge(item int) bool {
	count := 0 // 数字的位数
	for item != 0 {
		item /= 10
		count++
	}

	// fmt.Println(item, count)

	if count&1 == 1 { // count / 2 的余数
		return false
	}

	return true
}
