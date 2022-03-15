package _0220315

func countMaxOrSubsets(nums []int) int {
	//res := make([][]int, 0)
	res := 0
	maxSum := OrSum(nums)
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			if OrSum(path) == maxSum {
				res++
			}
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	//log.Println(res)

	return res
}

func OrSum(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res |= nums[i]
	}

	return res
}
