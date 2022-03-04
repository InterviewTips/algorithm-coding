package dp

func findLengthOfLCIS(nums []int) int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res[i] = res[i-1] + 1
		}
		if res[i] > result {
			result = res[i]
		}
	}
	//log.Println(res)

	return result
}
