package _0220220

import (
	"log"
	"sort"
)

// todo: 超时 92/115
func coutPairs(nums []int, k int) int64 {
	var count int64 = 0
	path := make([]int, 0)
	// 不确定是否有重复元素
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	memo := make(map[int]int)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == 2 {
			num1 := nums[path[0]]
			num2 := nums[path[1]]
			if memo[num1] == -1 && memo[num2] == -1 && memo[num1*num2] == -1 {
				return
			}
			if memo[num1] == 1 || memo[num2] == 1 || memo[num1*num2] == 1 {
				count++
				return
			}
			if num1%k == 0 {
				count++
				memo[num1] = 1
				return
			}
			memo[num1] = -1
			if num2%k == 0 {
				count++
				memo[num2] = 1
				return
			}
			memo[num2] = -1
			value := num1 * num2
			if value%k == 0 {
				count++
				memo[value] = 1
				return
			}
			// false
			memo[value] = -1
			return
		}

		for i := index; i < len(nums); i++ {
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	log.Println(count)

	return count
}
