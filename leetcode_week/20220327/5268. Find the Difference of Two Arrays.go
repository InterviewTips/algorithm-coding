package _0220327

// 求差集
func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = struct{}{}
	}
	nums2Map := make(map[int]struct{})
	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = struct{}{}
	}

	res1 := make([]int, 0)
	for k, _ := range nums1Map {
		if _, ok := nums2Map[k]; !ok {
			res1 = append(res1, k)
		}
	}

	res2 := make([]int, 0)
	for k, _ := range nums2Map {
		if _, ok := nums1Map[k]; !ok {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}
