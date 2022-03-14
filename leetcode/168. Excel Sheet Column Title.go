package leetcode

import "fmt"

// 10进制转26进制
func convertToTitle(columnNumber int) string {
	var res string
	for columnNumber != 0 {
		value := columnNumber % 26
		if value == 0 {
			res = "Z" + res
			columnNumber = columnNumber/26 - 1
			continue
		} else {
			res = fmt.Sprintf("%c%s", 'A'+value-1, res)
		}
		columnNumber /= 26
	}

	return res
}
