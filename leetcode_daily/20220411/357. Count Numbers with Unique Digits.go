package _0220411

import "log"

func countNumbersWithUniqueDigits(n int) int {

	data := make(map[int]int)
	data[0] = 1
	data[1] = 10
	var sub = func(num int) int {
		v, ok := data[num]
		if ok {
			return v
		}

		sum := 9
		for i := 0; i < n-1; i++ {
			sum *= 9 - i
		}
		log.Println(sum)
		data[num] = sum

		res := sum
		//for i := n-1; i >= 1; i-- {
		//	res += countNumbersWithUniqueDigits(i)
		//}
		res += countNumbersWithUniqueDigits(n - 1) // 只需要加下一个即可 因为下一个已经包含了先前所有情况

		return res
	}

	return sub(n)
}
