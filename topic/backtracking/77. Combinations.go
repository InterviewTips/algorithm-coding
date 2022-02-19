package backtracking

import "log"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtracking func(nVar int, pathVar []int)
	backtracking = func(nVar int, pathVar []int) {
		if len(pathVar) == k {
			value := make([]int, len(pathVar))
			copy(value, pathVar)
			res = append(res, value)
			return
		}

		for i := nVar; i <= n; i++ {
			pathVar = append(pathVar, i)
			backtracking(i+1, pathVar)
			pathVar = pathVar[:len(pathVar)-1]
		}
	}
	path := make([]int, 0)
	backtracking(1, path)

	log.Println(res)
	return res
}
