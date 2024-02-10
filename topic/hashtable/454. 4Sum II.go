package hashtable

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			numMap[nums1[i]+nums2[j]]++
		}
	}
	var count int
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			target := 0 - (nums3[i] + nums4[j])
			v, ok := numMap[target]
			if ok {
				count += v
			}
		}
	}

	return count
}
