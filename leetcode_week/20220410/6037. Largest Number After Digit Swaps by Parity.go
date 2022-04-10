package _0220410

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func largestInteger(num int) int {
	// 可以进行排序
	numStr := fmt.Sprintf("%d", num)
	num1 := make([]uint8, 0)
	num2 := make([]uint8, 0)
	for i := 0; i < len(numStr); i++ {
		item := numStr[i] - '0'
		if item&1 == 1 {
			num1 = append(num1, item)
		} else {
			num2 = append(num2, item)
		}
	}
	sort.SliceStable(num1, func(i, j int) bool {
		return num1[i] > num1[j]
	})
	sort.SliceStable(num2, func(i, j int) bool {
		return num2[i] > num2[j]
	})

	var res bytes.Buffer
	for i := 0; i < len(numStr); i++ {
		item := numStr[i] - '0'
		if item&1 == 1 {
			value := num1[0]
			num1 = num1[1:]
			res.WriteString(fmt.Sprintf("%d", value))
		} else {
			value := num2[0]
			num2 = num2[1:]
			res.WriteString(fmt.Sprintf("%d", value))
		}
	}

	ans, _ := strconv.Atoi(res.String())

	return ans
}
