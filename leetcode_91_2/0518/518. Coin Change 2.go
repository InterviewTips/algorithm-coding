package _518

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	coinsLength := len(coins)
	dp := make([][]int, coinsLength+1)
	for i := 0; i < coinsLength+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < coinsLength+1; i++ {
		dp[i][0] = 1 // amount 是 0 的话，需要初始化为 1
	}
	for i := 1; i < coinsLength+1; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 { // 可以放的情况下
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else { // 放不下
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[coinsLength][amount]
}

func change1(amount int, coins []int) int {
	coinsLength := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1 // base case
	for i := 1; i < coinsLength+1; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}
	return dp[amount]
}
