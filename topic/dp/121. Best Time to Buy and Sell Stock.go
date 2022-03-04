package dp

import "log"

func maxProfit(prices []int) int {
	// 只能买入一次
	res := make([]int, len(prices))
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			res[i] = res[i-1]
			continue
		}
		res[i] = getMax(res[i-1], prices[i]-min)
	}

	log.Println(res)

	return res[len(prices)-1]
}
