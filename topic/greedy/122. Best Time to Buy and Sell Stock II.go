package greedy

func maxProfit(prices []int) int {
	//profit := make([]int, 0)
	count := 0
	for i := 1; i < len(prices); i++ {
		if value := prices[i] - prices[i-1]; value > 0 {
			count += value
		}
	}

	return count
}
