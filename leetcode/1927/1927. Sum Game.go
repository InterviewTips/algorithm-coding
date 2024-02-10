package _1927

import "log"

// 参考题解: https://leetcode-cn.com/problems/sum-game/solution/shu-xue-jie-fa-by-newhar-fqjs/
func sumGame(num string) bool {
	sum1 := 0
	cnt1 := 0
	for i := 0; i < len(num)/2; i++ {
		if num[i] != '?' {
			sum1 += int(num[i] - '0')
		} else {
			cnt1++
		}
	}
	sum2 := 0
	cnt2 := 0
	for i := len(num) / 2; i < len(num); i++ {
		if num[i] != '?' {
			sum2 += int(num[i] - '0')
		} else {
			cnt2++
		}
	}

	log.Println(sum1, cnt1, sum2, cnt2)
	//delta := 9 * (cnt1 - cnt2) / 2 + sum1 - sum2 // 直接除以 2 会导致 floor 有问题
	delta := 9*(cnt1-cnt2) + 2*(sum1-sum2) // 改为乘以 2
	log.Println(delta)

	return delta != 0 // !=0 的时候 Alice 胜出，刚好返回 true
}
