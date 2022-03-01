package dp

import "log"

func numSquares(n int) int {
	res := make([]int, n+1)

	res[0] = 0 // 注意需要为 0
	//res[1] = 1
	for i := 1; i < len(res); i++ {
		res[i] = 10001
	}
	// 物品是啥呢
	for i := 1; i <= 100; i++ { // n 最多是 10000 所以最多也就选中 100 即可
		for j := i * i; j <= n; j++ {
			res[j] = getMin(res[j], res[j-i*i]+1) // 选中了 i 就 +1
		}
		log.Println("i is", i, res)
	}

	return res[n]
}
