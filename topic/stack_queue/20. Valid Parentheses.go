package stack_queue

/*
https://leetcode-cn.com/problems/valid-parentheses/
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

添加到 stack 中为右边对应即可，可以直接对比
1、左边剩余括号
2、右边剩余括号
3、符号不等
*/

func isValid(s string) bool {
	stack1 := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			stack1 = append(stack1, ']')
		case '{':
			stack1 = append(stack1, '}')
		case '(':
			stack1 = append(stack1, ')')
		default:
			if len(stack1) == 0 { // 右边剩余括号
				return false
			}
			// pop
			popRes := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			if popRes != s[i] { // 符号不等
				return false
			}
		}
	}
	if len(stack1) > 0 {
		return false
	}

	return true
}
