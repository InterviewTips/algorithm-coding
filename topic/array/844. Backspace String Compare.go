package array

func backspaceCompare(s string, t string) bool {
	return backspaceCompare1(s, t)
}

func backspaceCompare1(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	skipS := 0
	skipT := 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else { // 没有退格 正常 break 进行字符的比较
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		// 处理其他情况
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}
