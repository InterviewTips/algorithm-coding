package _748

func sumOfUnique(nums []int) int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		data[nums[i]]++
	}

	sum := 0
	for k, v := range data {
		if v == 1 {
			sum += k
		}
	}

	return sum
}
