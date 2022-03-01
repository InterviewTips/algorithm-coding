package dp

func change(amount int, coins []int) int {
	res := make([]int, amount+1)
	// res[j] amount = j 有多少种方式
	res[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			res[j] += res[j-coins[i]]
		}
		//log.Println(res)
	}

	return res[amount]
}
