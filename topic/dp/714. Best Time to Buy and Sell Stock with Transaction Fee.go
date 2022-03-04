package dp

// var maxProfit = maxProfit6
func maxProfit6(prices []int, fee int) int {
	res := make([][2]int, len(prices))

	res[0][0] = -prices[0]
	res[0][1] = 0

	for i := 1; i < len(prices); i++ {
		res[i][0] = getMax(res[i-1][0], res[i-1][1]-prices[i])
		res[i][1] = getMax(res[i-1][1], res[i-1][0]+prices[i]-fee)
	}

	return res[len(prices)-1][1]
}
