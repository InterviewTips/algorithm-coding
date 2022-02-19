package backtracking

func letterCombinations(digits string) []string {
	data := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)
	if len(digits) == 0 { // 空处理
		return res
	}
	path := make([]byte, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		// len(digits) 为组成的 path 需要的长度
		if len(path) == len(digits) { // 也可以是 index == len(digits)
			res = append(res, string(path))
			return
		}

		// 取出数字对应的字母列表
		letters := data[digits[index]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtracking(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return res
}
