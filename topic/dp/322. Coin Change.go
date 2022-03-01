package dp

import (
	"log"
)

func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	res[0] = 0
	// 关键点: 初始化
	for i := 1; i < len(res); i++ {
		res[i] = 10001
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			res[j] = getMin(res[j], res[j-coins[i]]+1)
		}
		log.Println(res)
	}

	if res[amount] == 10001 {
		return -1
	}

	return res[amount]
}
