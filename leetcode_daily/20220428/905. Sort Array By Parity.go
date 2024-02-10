package _0220428

func sortArrayByParity(nums []int) []int {
	ans1 := make([]int, 0)
	ans2 := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			ans1 = append(ans1, nums[i])
		} else {
			ans2 = append(ans2, nums[i])
		}
	}

	return append(ans2, ans1...)
}
