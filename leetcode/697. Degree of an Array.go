package leetcode

func findShortestSubArray(nums []int) int {
	data := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		item := nums[i]
		_, ok := data[item]
		if !ok {
			data[item] = []int{i}
			continue
		}

		data[item] = append(data[item], i)
	}

	maxCount := 0
	for _, v := range data {
		if len(v) > maxCount {
			maxCount = len(v)
		}
	}

	minCount := 50000
	for _, v := range data {
		if len(v) == maxCount {
			value := v[len(v)-1] - v[0] + 1
			if value < minCount {
				minCount = value
			}
		}
	}

	return minCount
}
