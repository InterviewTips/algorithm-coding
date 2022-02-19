package backtracking

import "log"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(nVar int, sum int)
	backtracking = func(nVar int, sum int) {
		if sum > n { // 剪枝
			return
		}
		if len(path) == k {
			if sum == n {
				log.Println("now path is ", path)
				value := make([]int, len(path))
				copy(value, path)
				res = append(res, value)
			}
			return
		}

		// 剪枝
		for i := nVar; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			//log.Println("path is", path, "sum is", sum)
			backtracking(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtracking(1, 0)

	log.Println(res)
	return res
}
