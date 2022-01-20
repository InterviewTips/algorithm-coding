package two_pointers

func replaceSpace(s string) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		}
	}
	res := make([]byte, len(s)+count*2)
	index := len(s) - 1 // 指向旧的末尾
	for j := len(res) - 1; j >= 0; j-- {
		if s[index] != ' ' {
			res[j] = s[index]
		} else {
			res[j] = '0'
			res[j-1] = '2'
			res[j-2] = '%'
			j = j - 2 // 等一下还会 --
		}

		index--
		//if index < 0 {// 理论上不需要
		//	break
		//}
	}

	return string(res)
}
