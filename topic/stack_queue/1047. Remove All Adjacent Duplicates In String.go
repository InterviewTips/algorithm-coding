package stack_queue

/*栈的应用
输入："abbaca"
输出："ca"
*/

func removeDuplicates(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i > 0 && len(res) > 0 && res[len(res)-1] == s[i] {
			// 回退
			res = res[:len(res)-1]
			continue
		}

		res = append(res, s[i])
	}

	return string(res)
}
