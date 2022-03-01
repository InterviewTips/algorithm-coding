package dp

func climbStairsCompletePack(n int, m int) int {
	// 爬楼梯完全背包写法
	// m 表示一次可以爬 [1-m] 层台阶
	// res[i] 表示爬 i 层台阶有多少种方式
	res := make([]int, n+1)
	res[0] = 1 // 爬 0 层的方式就是不爬
	// 排列模式 物品在里层 容量在外层
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i-j >= 0 {
				res[i] += res[i-j]
			}
		}
	}

	return res[n]
}
