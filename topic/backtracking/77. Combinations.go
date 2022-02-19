package backtracking

import "log"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(nVar int)
	backtracking = func(nVar int) {
		if len(path) == k {
			value := make([]int, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		// 剪枝
		//for i := nVar; i <= n; i++ {
		// 已经选择的元素长度 len(path)
		// k - len(path) 还需要选择的元素长度
		// n - i + 1 >= k - len(path) 即可 表示 从 i 到 n 的元素个数 >= 还需选择的元素长度
		for i := nVar; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)

	log.Println(res)
	return res
}
