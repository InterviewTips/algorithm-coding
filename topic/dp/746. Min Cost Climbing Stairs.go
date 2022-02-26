package dp

func minCostClimbingStairs(cost []int) int {
	// 题意 cost.len >= 2
	if len(cost) == 2 {
		return getMin(cost[0], cost[1]) // 不需要花费
	}

	// 注意：楼顶是比最后一个元素还后面
	// res[i] 表示爬到第 i 层需要花费的 sum
	// res[i] = getMin(res[i-1]+cost[i-1], res[i-2]+cost[i-2])
	// 爬到 i-1 需要花费的 sum + cost[i-1] 再爬一层可以爬到 i
	// 爬到 i-2 需要花费的 sum + cost[i-2] 再爬两层可以爬到 i
	// 取出两者最小值
	res := make([]int, len(cost)+1)
	res[0] = 0
	res[1] = 0
	for i := 2; i <= len(cost); i++ {
		res[i] = getMin(res[i-1]+cost[i-1], res[i-2]+cost[i-2])
	}

	return res[len(cost)]
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
