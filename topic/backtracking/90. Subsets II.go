package backtracking

import (
	"log"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// 排序 才可以排除前后相同的元素
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	res = append(res, []int{})
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] { // 到达下一个树枝 横向
				continue
			}
			path = append(path, nums[i])
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			backtracking(i + 1) // 纵向
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	log.Println(res)
	return res
}
