package backtracking

import (
	"log"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var backtracking func(numsVar []int)
	backtracking = func(numsVar []int) {
		if len(path) == len(nums) {
			//log.Println("path is", path)
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := 0; i < len(numsVar); i++ {
			if i > 0 && numsVar[i] == numsVar[i-1] { // 跳过同一层
				continue
			}
			path = append(path, numsVar[i])
			// 去除 nums i
			newNumsVar := make([]int, len(numsVar[:i]))
			copy(newNumsVar, numsVar[:i])
			if i+1 <= len(numsVar) {
				newNumsVar = append(newNumsVar, numsVar[i+1:]...)
			}
			backtracking(newNumsVar)
			path = path[:len(path)-1]
		}

	}

	backtracking(nums)

	log.Println(res)
	return res
}
