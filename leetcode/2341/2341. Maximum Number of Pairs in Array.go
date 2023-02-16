package _341

func numberOfPairs(nums []int) []int {
	data := make(map[int]struct{})
	pairs := 0
	count := len(nums)
	for _, item := range nums {
		_, ok := data[item]
		if ok {
			pairs++
			count -= 2
			delete(data, item)
			continue
		}
		data[item] = struct{}{}
	}

	return []int{pairs, count}
}
