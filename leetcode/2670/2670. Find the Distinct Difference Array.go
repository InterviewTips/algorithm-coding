package _670

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	diff := make([]int, n)
	prefixSet := make(map[int]bool)
	suffixSet := make(map[int]int)

	// Initialize suffix set with counts
	for _, num := range nums {
		suffixSet[num]++
	}

	for i, num := range nums {
		prefixSet[num] = true
		suffixSet[num]--
		if suffixSet[num] == 0 {
			delete(suffixSet, num)
		}

		diff[i] = len(prefixSet) - len(suffixSet)
	}
	return diff
}
