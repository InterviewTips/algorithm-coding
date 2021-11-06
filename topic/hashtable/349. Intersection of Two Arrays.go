package hashtable

func intersection(nums1 []int, nums2 []int) []int {
	// hash table 一个 map 即可
	nums1Map := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = struct{}{} // 设置标识位即可
	}

	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		// 这里要考虑 nums2 中可能有重复的值，所以如果
		// nums1 中有了, append 之后要从 map 中删除
		_, ok := nums1Map[nums2[i]]
		if ok {
			res = append(res, nums2[i])
			delete(nums1Map, nums2[i]) // 避免 append 重复的值
		}
	}

	return res
}
