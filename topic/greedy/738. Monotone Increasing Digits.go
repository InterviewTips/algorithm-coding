package greedy

import "strconv"

func monotoneIncreasingDigits(n int) int {
	strNum := []byte(strconv.Itoa(n))
	niceFlag := len(strNum)
	for i := len(strNum) - 1; i > 0; i-- {
		if strNum[i-1] > strNum[i] {
			niceFlag = i
			strNum[i-1]--
		}
	}
	for i := niceFlag; i < len(strNum); i++ {
		strNum[i] = '9'
	}

	value, _ := strconv.Atoi(string(strNum))

	return value
}
