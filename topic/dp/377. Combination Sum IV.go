package dp

func combinationSum4(nums []int, target int) int {
	res := make([]int, target+1)
	res[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				res[j] += res[j-nums[i]]
			}
		}
		//log.Println(res)
	}

	return res[target]
}
