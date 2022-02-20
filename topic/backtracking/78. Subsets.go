package backtracking

import "log"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	// 加入空集合
	res = append(res, []int{})
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	log.Println(res)

	return res
}
