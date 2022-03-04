package dp

// var maxProfit = maxProfit4
func maxProfit4(k int, prices []int) int {

	// 边界情况 0 <= prices.length <= 1000
	if len(prices) == 0 {
		return 0
	}

	// 2*k+1
	res := make([][]int, len(prices))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, 2*k+1)
	}
	for i := 0; i < 2*k+1; i++ {
		if i&1 == 1 {
			res[0][i] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		res[i][0] = res[i-1][0]
		for j := 1; j <= 2*k; j++ {
			if j&1 == 1 {
				res[i][j] = getMax(res[i-1][j], res[i-1][j-1]-prices[i]) // 买入
			} else {
				res[i][j] = getMax(res[i-1][j], res[i-1][j-1]+prices[i]) // 卖出
			}
		}
	}

	return res[len(prices)-1][2*k]
}
