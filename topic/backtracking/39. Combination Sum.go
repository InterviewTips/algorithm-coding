package backtracking

import (
	"log"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	// 排序 从小到大
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
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

		// 剪枝
		// candidates 已经排序
		// sum + candidates[i] > target 则后面的就不用比较了 自然是更大
		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			backtracking(sum+candidates[i], i) // 关键: 传入 i 因为每次都可以从 i 开始
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)

	log.Println(res)

	return res
}
