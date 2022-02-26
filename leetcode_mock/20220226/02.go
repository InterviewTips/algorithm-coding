package _0220226

// https://leetcode-cn.com/problems/increasing-subsequences/

import "log"

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == len(nums) {
			return
		}

		used := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			if len(path) >= 2 {
				value := make([]int, len(path))
				copy(value, path)
				res = append(res, value)
			}
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	log.Println(res)

	return res
}
