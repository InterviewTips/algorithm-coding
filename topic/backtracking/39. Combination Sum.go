package backtracking

import (
	"log"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(sum int, index int)
	backtracking = func(sum int, index int) {
		if sum > target { // 剪枝
			return
		}

		if sum == target {
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtracking(sum+candidates[i], i) // 关键: 传入 i 因为每次都可以从 i 开始
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)

	log.Println(res)

	return res
}
