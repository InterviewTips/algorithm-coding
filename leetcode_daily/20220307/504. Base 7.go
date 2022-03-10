package _0220307

import (
	"fmt"
	"strings"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	flag := false
	if num < 0 { // 负数先转为正数然后最后再加 - flag
		flag = true
		num = -num
	}

	res := make([]string, 0)
	for num != 0 {
		value := num % 7
		res = append(res, fmt.Sprintf("%v", value))
		num = num / 7
	}

	value := reverse(res)
	if flag {
		value = fmt.Sprintf("-%s", value)
	}

	return value
}

func reverse(res []string) string {
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return strings.Join(res, "")
}
