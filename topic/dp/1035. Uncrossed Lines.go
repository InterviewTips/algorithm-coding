package dp

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// 最长公共子序列
	res := make([][]int, len(nums1)+1)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = getMax(res[i][j-1], res[i-1][j])
			}
		}
	}

	return res[len(nums1)][len(nums2)]
}
