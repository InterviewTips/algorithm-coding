package theory

import "log"

func twoDimensional01Bag(weight, value []int, bagWeight int) int {
	// res[i][j] 表示 从 0~i 物品，取出物品存放到最大容量为 j 的背包,使得其价值最大的总和
	// res[i][j] = getMax(res[i-1][j], res[i-1][j-weight[i]] + value[i])
	// res[i-1][j] 背包容量已经是 j 的最大总和 放不下 物品i
	// res[i-1][j-weight[i]] + value[i] 放的下物品 i
	res := make([][]int, len(weight))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, bagWeight+1)
	}
	// 初始化第一行
	for j := weight[0]; j <= bagWeight; j++ {
		res[0][j] = value[0]
	}
	for i := 1; i < len(weight); i++ { // 遍历物品
		for j := 1; j <= bagWeight; j++ { // 如果背包容量为 0，其实 res[i][0] = 0
			// 注意边界
			if j-weight[i] >= 0 {
				res[i][j] = getMax(res[i-1][j], res[i-1][j-weight[i]]+value[i])
			} else {
				res[i][j] = res[i-1][j]
			}
		}
	}

	log.Println(res)

	return res[len(weight)-1][bagWeight]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
