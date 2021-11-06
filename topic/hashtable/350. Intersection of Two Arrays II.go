package hashtable

func intersect(nums1 []int, nums2 []int) []int {
	// 两个 map 都要存下来
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]]++
	}

	res := make([]int, 0)
	for k, v1 := range nums1Map {
		v2, ok := nums2Map[k]
		if ok {
			if v1 >= v2 { // v2
				for i := 0; i < v2; i++ {
					res = append(res, k)
				}
			} else { // v1
				for i := 0; i < v1; i++ {
					res = append(res, k)
				}
			}
		}
	}

	return res
}
