package backtracking

import (
	"log"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(candidates))
	// 排序 从小到大
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var backtracking func(sum int, index int)
	backtracking = func(sum int, index int) {
		if sum > target {
			return
		}

		if sum == target {
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				// 表示同一个树层 前面已经选取过
				continue
			}
			used[i] = true
			path = append(path, candidates[i])
			backtracking(sum+candidates[i], i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)

	log.Println(res)

	return res
}

// var combinationSum2 = combinationSum2Another
func combinationSum2Another(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	// 排序 从小到大
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var backtracking func(sum int, index int)
	backtracking = func(sum int, index int) {
		if sum > target {
			return
		}

		if sum == target {
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			// i > index 表示到了下一个树枝
			if i > index && candidates[i] == candidates[i-1] {
				// 表示同一个树层 前面已经选取过
				continue
			}
			path = append(path, candidates[i])
			backtracking(sum+candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)

	log.Println(res)

	return res
}
