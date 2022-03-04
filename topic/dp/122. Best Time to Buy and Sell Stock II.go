package dp

import "log"

// var maxProfit = maxProfit2
func maxProfit2(prices []int) int {
	// 可以买入多次
	res := make([]int, len(prices))
	res[0] = 0
	for i := 1; i < len(prices); i++ {
		res[i] = getMax(res[i-1], res[i-1]+prices[i]-prices[i-1])
	}

	log.Println(res)

	return res[len(prices)-1]
}
