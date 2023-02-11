package _217

func containsDuplicate(nums []int) bool {
	data := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		_, ok := data[nums[i]]
		if !ok {
			data[nums[i]] = struct{}{}
			continue
		}
		return true
	}
	return false
}
