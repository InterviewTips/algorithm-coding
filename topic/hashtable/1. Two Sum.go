package hashtable

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		v, ok := numMap[t]
		if !ok {
			numMap[nums[i]] = i
			continue
		}
		return []int{v, i}
	}

	return []int{}
}
