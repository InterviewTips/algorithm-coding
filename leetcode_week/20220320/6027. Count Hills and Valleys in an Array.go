package _0220320

func countHillValley(nums []int) int {
	// 先去除重复的
	slowIndex := 0
	dataNums := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != dataNums[slowIndex] {
			dataNums = append(dataNums, nums[i])
			slowIndex++
		}
	}
	//log.Println("nums is ", dataNums)
	// 计算峰和谷的数量
	count := 0
	for i := 0; i < len(dataNums); i++ {
		if i == 0 || i == len(dataNums)-1 {
			continue
		}
		// 峰 或者 谷
		if (dataNums[i-1] < dataNums[i] && dataNums[i] > dataNums[i+1]) ||
			(dataNums[i-1] > dataNums[i] && dataNums[i] < dataNums[i+1]) {
			count++
		}
	}

	return count
}
