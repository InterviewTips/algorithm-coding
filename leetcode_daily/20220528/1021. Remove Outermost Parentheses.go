package _0220528

import "bytes"

func removeOuterParentheses(req string) string {
	if len(req) == 0 {
		return req
	}
	stack := make([]byte, 0)
	var res bytes.Buffer
	var start int
	for i := 0; i < len(req); i++ {
		item := req[i]
		switch item {
		case '(':
			stack = append(stack, item)
		case ')':
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res.WriteString(req[start+1 : i])
				start = i + 1 // 开始位置在下一个 ( 括号
			}
		}
	}

	return res.String()
}
