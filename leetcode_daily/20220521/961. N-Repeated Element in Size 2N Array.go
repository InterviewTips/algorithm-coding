package _0220521

func repeatedNTimes(nums []int) int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := data[nums[i]]
		if ok {
			return nums[i]
		}
		data[nums[i]]++
	}

	return -1
}