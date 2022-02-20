package backtracking

import "log"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
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
			path = append(path, numsVar[i])

			//newNumsVar := numsVar[:i] 错误写法 需要使用 copy
			// 去除 nums i
			newNumsVar := make([]int, len(numsVar[:i]))
			copy(newNumsVar, numsVar[:i])
			if i+1 <= len(numsVar) {
				// todo: 也可以使用 used 记录之前是否已经选中
				// 上面如果直接使用 newNumsVar := numsVar[:i] 地址和 nums 是相同的，有可能导致 nums 值被覆盖
				// 2022/02/20 15:45:43 new cap: 3, new pointer: 0x1400001a108, old 0x1400001a108
				//log.Printf("new cap: %v, new pointer: %p, old %p",cap(newNumsVar), newNumsVar, nums)
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
