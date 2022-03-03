package dp

// var maxProfit = maxProfit5
func maxProfit5(prices []int) int {
	res := make([][4]int, len(prices))

	res[0][0] = -prices[0] // 买入股票状态
	res[0][1] = 0          // 两天前就卖出了股票，度过了冷冻期，一直没操作，今天保持卖出股票状态
	res[0][2] = 0          // 当天卖出了股票
	res[0][3] = 0          // 冷冻期

	for i := 1; i < len(prices); i++ {
		res[i][0] = getMax(res[i-1][0], getMax(res[i-1][3]-prices[i], res[i-1][1]-prices[i]))
		res[i][1] = getMax(res[i-1][1], res[i-1][3])
		res[i][2] = res[i-1][0] + prices[i]
		res[i][3] = res[i-1][2]
	}

	value := res[len(prices)-1]
	return getMax3(value[1], value[2], value[3])
}
