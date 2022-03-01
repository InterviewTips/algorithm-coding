package theory

func oneDimensionalCompletePack(weight, value []int, bagWeight int) int {
	// 完全背包，每个物品可以取多次
	res := make([]int, bagWeight+1)
	res[0] = 0

	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			res[j] = getMax(res[j], res[j-weight[i]]+value[i])
		}
	}

	return res[bagWeight]
}
