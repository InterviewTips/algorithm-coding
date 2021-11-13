package hashtable

import (
	"log"
)

func isHappy(n int) bool {
	// 通过 map 记录出现过的 sum
	sumMap := make(map[int]struct{})
	for {
		res := countNum(n)
		log.Println(res)
		if res == 1 {
			return true
		}
		_, ok := sumMap[res]
		if ok {
			return false
		} else {
			sumMap[res] = struct{}{}
			n = res
		}
		//time.Sleep(1*time.Second)
	}
}

func countNum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10) // 平方
		n /= 10
	}
	return sum
}
