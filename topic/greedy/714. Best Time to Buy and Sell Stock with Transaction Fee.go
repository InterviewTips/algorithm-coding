package greedy

// var maxProfit = maxProfitFee
func maxProfitFee(prices []int, fee int) int {
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] > minPrice+fee {
			profit += prices[i] - minPrice - fee // 收益
			minPrice = prices[i] - fee           // 减掉 fee 如果不是上升区间收益的最后一天 下次 -minPrice 可以将 fee 抵消掉 无需重复扣除
		}
	}

	return profit
}
