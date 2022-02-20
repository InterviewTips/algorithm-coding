package backtracking

import (
	"log"
)

// 至少要两个元素
func findSubsequences(nums []int) [][]int {
	// 不应该排序 因为题目要求有序
	//sort.SliceStable(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	used := make([]bool, len(nums))
	path := make([]int, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			if i > index && isNumExists(index, i, nums) {
				// todo: 也可以函数 backtracking 定义一个 map 去重
				// 判断条件挺关键的 因为没有排序 所以要判断 nums[index,i) 是否有跟 nums[i] 相等的
				// 有说明同个树层已经选过了相同的数字了
				continue
			}

			if len(path) >= 1 {
				if nums[i] >= path[len(path)-1] { // 递增
					path = append(path, nums[i])
					//log.Println("path is", path)
					value := make([]int, len(path))
					copy(value, path)
					res = append(res, value)
				} else {
					continue
				}
			} else { // len.path <= 1 还没有元素 直接添加
				path = append(path, nums[i])
			}
			used[i] = true
			backtracking(i + 1)
			//log.Println("now path is",path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtracking(0)

	log.Println(res)

	return res
}

func isNumExists(start int, end int, nums []int) bool {
	for i := start; i < end; i++ {
		if nums[i] == nums[end] {
			return true
		}
	}

	return false
}
