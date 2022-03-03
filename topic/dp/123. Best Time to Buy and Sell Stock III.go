package dp

// var maxProfit = maxProfit3
func maxProfit3(prices []int) int {
	res := make([][5]int, len(prices))
	res[0][0] = 0
	res[0][1] = -prices[0]
	res[0][2] = 0
	res[0][3] = -prices[0]
	res[0][4] = 0

	for i := 1; i < len(prices); i++ {
		res[i][0] = res[i-1][0]
		res[i][1] = getMax(res[i-1][1], res[i-1][0]-prices[i])
		res[i][2] = getMax(res[i-1][2], res[i-1][1]+prices[i])
		res[i][3] = getMax(res[i-1][3], res[i-1][2]-prices[i])
		res[i][4] = getMax(res[i-1][4], res[i-1][3]+prices[i])
	}

	return res[len(prices)-1][4]
}
