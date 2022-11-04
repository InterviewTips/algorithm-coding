package issue9

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	ptr := 0
	var stk []string

	var getDigits func() string
	getDigits = func() string {
		// 获取数字
		ret := ""
		for ; s[ptr] >= '0' && s[ptr] <= '9' && ptr < len(s); ptr++ {
			ret += string(s[ptr])
		}
		return ret
	}

	var getString func(v []string) string
	getString = func(v []string) string {
		return strings.Join(v, "")
	}

	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' { // 获取数字
			digits := getDigits()
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else { // 右花括号
			var sub []string
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1] // 出栈
			}
			stk = stk[:len(stk)-1] // 将 "[" pop 出栈
			// 逆序
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1] // 将数字 pop 出栈
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t) // 重新入栈
			ptr++
		}
	}

	return getString(stk)
}
