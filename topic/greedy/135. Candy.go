package greedy

import "log"

func candy(ratings []int) int {
	res := make([]int, len(ratings))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	// 从前到后
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	log.Println(res)
	// 从后到前
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = getMax(res[i], res[i+1]+1)
		}
	}
	log.Println(res)

	return getSum(res)
}
